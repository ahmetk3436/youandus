package repository

import (
	"errors"
	"gorm.io/gorm"
	"sync"
	"youandus/internal/storage"
)

type VerificationRepoInterface interface {
	GetEmailVerificationCode(email, verificationCode string)
	CheckSMSVerificationCode(phoneNumber, code string) (bool, error)
	CheckEmailVerificationCode(email, code string) (bool, error)
}

type VerificationRepo struct {
	db *gorm.DB
}

func NewVerificationRepo(db *gorm.DB) *VerificationRepo {
	return &VerificationRepo{
		db: db,
	}
}

func (c *VerificationRepo) CheckSMSVerificationCode(phoneNumber, code string) (bool, error) {
	// SMS doğrulama kodunu kontrol etme işlemini gerçekleştir
	// TODO: SMS doğrulama kodu kontrolünü yap ve sonucu dön

	return false, nil
}

func (c *VerificationRepo) CheckEmailVerificationCode(email, code string) (bool, uint, error) {
	verificationCode := GetEmailVerificationCode(email)
	if verificationCode != code {
		return false, 0, errors.New("verification code is false")
	}
	if err := c.db.Table("user_registers").Where("user_email = ?", email).Update("email_verified", 1).Error; err != nil {
		return false, 0, err
	}
	var userID uint
	if err := c.db.Table("user_registers").Where("user_email = ?", email).Pluck("id", &userID).Error; err != nil {
		return false, 0, err
	}
	return true, userID, nil
}
func GetEmailVerificationCode(email string) string {
	db := storage.GetDB()
	var verificationCode string
	var wg sync.WaitGroup
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = db.Table("user_registers").Where("user_email = ?", email).Pluck("verification_code", &verificationCode).Error
	}()
	wg.Wait()
	if err != nil {
		return ""
	}
	return verificationCode
}
