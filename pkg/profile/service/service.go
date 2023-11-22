package service

import (
	"errors"
	"youandus/pkg/profile/model"
	"youandus/pkg/profile/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateProfile(userID uint) error {
	err := s.repo.CreateProfile(userID)
	if err != nil {
		return errors.New("kullanıcının profili oluşturulamadı ")
	}
	return nil
}
func (s Service) GetProfile(userID uint) (model.User, error) {
	profile, err := s.repo.GetProfile(userID)
	if err != nil {
		return model.User{}, errors.New("kullanıcının profili oluşturulamadı ")
	}
	return profile, nil
}
func (s Service) UpdateProfile(userID uint, newProfile model.User) (profile model.User, err error) {
	if userID <= 0 {
		return model.User{}, errors.New("kullanıcı bulunamadı ")
	}
	profile, err = s.repo.UpdateProfile(userID, newProfile)
	if err != nil {
		return model.User{}, errors.New("profil güncellenemedi ")
	}
	return profile, nil
}
func (s Service) DeleteProfile(userID uint) error {
	err := s.repo.DeleteProfile(userID)
	if err != nil {
		return errors.New("profil silinemedi !" + err.Error())
	}
	return nil
}
