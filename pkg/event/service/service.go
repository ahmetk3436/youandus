package service

import (
	"errors"
	"youandus/pkg/event/model"
	"youandus/pkg/event/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateEvent(event *model.Event) error {
	if err := s.repo.CreateEvent(event); err != nil {
		return errors.New("Etkinlik oluşturulamadı")
	}
	return nil
}

func (s Service) GetEvent(eventID, userID uint) (*model.Event, error) {
	eventData, err := s.repo.GetEvent(eventID, userID)
	if err != nil {
		return nil, errors.New("Etkinlik bulunamadı")
	}
	return eventData, nil
}

func (s Service) GetEvents(userID uint) ([]*model.Event, error) {
	eventDatas, err := s.repo.GetEvents(userID)
	if err != nil {
		return nil, errors.New("Etkinlik bulunamadı")
	}
	return eventDatas, nil
}

func (s Service) UpdateEvent(eventID uint, newEvent model.Event) (*model.Event, error) {
	if eventID <= 0 {
		return nil, errors.New("Etkinlik bulunamadı")
	}

	eventData, err := s.repo.UpdateEvent(eventID, newEvent)
	if err != nil {
		return nil, errors.New("Etkinlik güncellenemedi")
	}

	return eventData, nil
}

func (s Service) DeleteEvent(eventID uint) error {
	if err := s.repo.DeleteEvent(eventID); err != nil {
		return errors.New("Etkinlik silinemedi: " + err.Error())
	}
	return nil
}
