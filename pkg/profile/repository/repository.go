package repository

import (
	"errors"
	"youandus/pkg/profile/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) UpdateProfile(userID uint, profile model.User) (model.User, error) {
	profile.UserID = userID

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&model.User{}, "user_id = ?", userID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			if err := tx.Create(&profile).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&model.User{}).Where("user_id = ?", userID).Updates(&profile).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return model.User{}, err
	}

	var newProfile model.User
	if err := r.db.First(&newProfile, "user_id = ?", userID).Error; err != nil {
		return model.User{}, err
	}

	return newProfile, nil
}

func (r Repository) GetProfile(userID uint) (model.User, error) {
	var profile model.User
	if err := r.db.Table("users").Where("user_id = ?", userID).Find(&profile).Error; err != nil {
		return model.User{}, err
	}
	return profile, nil
}
func (r Repository) DeleteProfile(userID uint) error {
	if err := r.db.Where("user_id = ?", userID).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
func (r Repository) CreateProfile(userID uint) error {
	profile := model.User{UserID: userID}
	if err := r.db.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}
