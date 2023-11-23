package repository

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"
	"youandus/internal/storage"
	"youandus/pkg/users/model"
	user2 "youandus/pkg/users/model/user"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	usersTable = "user_registers"
)

type UsersRepository struct {
	DB       *gorm.DB
	Redis    *storage.RedisClient
	RabbitMQ *amqp.Channel
}

func NewUsersRepository(db *gorm.DB, redis *storage.RedisClient, rabbitMQ *amqp.Channel) *UsersRepository {
	return &UsersRepository{
		DB:       db,
		Redis:    redis,
		RabbitMQ: rabbitMQ,
	}
}

// for admin user
func (u *UsersRepository) GetUsers() (*[]user2.UserRegister, error) {
	var users []user2.UserRegister
	err := u.DB.Table(usersTable).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *UsersRepository) GetUser(id int) (*user2.UserRegister, error) {
	var user user2.UserRegister
	err := u.DB.Table(usersTable).First(&user, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user, nil
}

func (u *UsersRepository) CreateUser(user *user2.UserRegister) (*model.BaseResponse, error) {
	hashedPass, err := HashPassword(user.UserPass)
	if err != nil {
		return nil, errors.New(err.Error() + " - hash")
	}
	user.UserPass = hashedPass
	verificationCode, err := GenerateVerificationCode(20)
	if err != nil {
		return &model.BaseResponse{}, err
	}
	user.VerificationCode = verificationCode
	if err = u.DB.Create(&user).Error; err != nil {
		return nil, errors.New(err.Error() + " - create")
	}
	user.UserPass = ""
	baseResponse := model.BaseResponse{
		Message: "User Created Successfully !",
		Data:    user,
	}
	userData, err := EncodeUserData(user.ID, user.UserEmail)
	if err != nil {
		return &model.BaseResponse{}, err
	}
	storage.PublishEmailVerification(userData)
	return &baseResponse, nil
}
func GenerateVerificationCode(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	verificationCode := base64.RawURLEncoding.EncodeToString(buffer)[:length]
	return verificationCode, nil
}
func EncodeUserData(userID uint, userEmail string) ([]byte, error) {
	user := struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}{
		ID:    userID,
		Email: userEmail,
	}

	encodedData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	return encodedData, nil
}
func (u *UsersRepository) LoginUser(login *user2.UserLogin) (*model.BaseResponse, error) {
	if login.Email == "" || login.Password == "" {
		return nil, errors.New("please provide email and password")
	}
	var loginedUser *user2.UserRegister
	err := u.DB.Table("user_registers").Where("user_email = ?", login.Email).First(&loginedUser).Error
	if err != nil {
		return nil, errors.New("error finding user")
	}

	if !CheckPasswordHash(login.Password, loginedUser.UserPass) {
		return nil, errors.New("invalid login credentials")
	}
	if !loginedUser.EmailVerified || !loginedUser.SmsVerified {
		return nil, errors.New("this user is not verified")
	}
	token, err := CreateToken(time.Hour*72, loginedUser.UserName, loginedUser.ID)
	if err != nil {
		return nil, errors.New("error creating token")
	}

	if err := u.DB.Table("user_registers").Where("user_email = ?", login.Email).Update("token", token).Error; err != nil {
		return nil, errors.New("error updating user token")
	}

	baseResponse := model.BaseResponse{
		Message: "Logined Successfully !",
		Data: &user2.UserLoginDTO{
			Email:    login.Email,
			UserName: loginedUser.UserName,
			Token:    *token,
		},
	}
	return &baseResponse, nil
}

func (u *UsersRepository) UpdateUser(id uint, user *user2.UserRegister) (*model.BaseResponse, error) {
	oldUser, err := u.GetUser(int(id))
	if err != nil {
		return nil, err
	}
	if user.UserEmail != "" {
		oldUser.UserEmail = user.UserEmail
	}
	if user.UserName != "" {
		oldUser.UserName = user.UserName
	}
	if user.UserName == "" && user.UserEmail == "" {
		return nil, errors.New("User can't updated")
	}
	err = u.DB.Where("id = ?", id).Updates(&oldUser).Error
	if err != nil {
		return nil, err
	}
	baseResponse := model.BaseResponse{
		Message: "User updated successfully !",
		Data:    oldUser,
	}
	return &baseResponse, nil
}

func (u *UsersRepository) DeleteUser(id uint) error {
	err := u.DB.Table(usersTable).Where("id = ?", id).Delete(&user2.UserRegister{}).Error
	if err != nil {
		return err
	}
	return nil
}
