package user

import "gorm.io/gorm"

type UserDTO struct {
	ID       uint
	UserMail string `gorm:"not null" json:"userMail"`
	UserName string `json:"userName" gorm:"unique"`
	Token    string `json:"token,omitempty"`
}
type UserRegister struct {
	gorm.Model
	UserPass         string `json:"password"`
	UserEmail        string `json:"userEmail" gorm:"unique"`
	UserName         string `json:"userName" gorm:"unique"`
	VerificationCode string
	EmailVerified    bool   `gorm:"default:0"`
	SmsVerified      bool   `gorm:"default:1"`
	Token            string `json:"token,omitempty"`
}
type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UserRegister) ToDTOCreate() *UserDTOResponse {
	return &UserDTOResponse{
		Message: "User created successfully",
		User: UserDTO{
			ID:       u.ID,
			UserMail: u.UserEmail,
			UserName: u.UserName,
		},
	}
}

func (u *UserRegister) ToDTOLogin() *UserDTOResponse {
	return &UserDTOResponse{
		Message: "User login succesfully !",
		User: UserDTO{
			ID:       u.ID,
			UserName: u.UserName,
			UserMail: u.UserEmail,
			Token:    u.Token,
		},
	}
}
