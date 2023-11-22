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
	err := s.repo.CreateEvent(event)
	if err != nil {
		return errors.New("etkinlik oluşturulamadı")
	}
	return nil
}

func (s Service) GetEvent(eventID uint) (model.Event, error) {
	eventData, err := s.repo.GetEvent(eventID)
	if err != nil {
		return model.Event{}, errors.New("etkinlik bulunamadı")
	}
	return eventData, nil
}

func (s Service) UpdateEvent(eventID uint, newEvent model.Event) (eventData model.Event, err error) {
	if eventID <= 0 {
		return model.Event{}, errors.New("etkinlik bulunamadı")
	}
	eventData, err = s.repo.UpdateEvent(eventID, newEvent)
	if err != nil {
		return model.Event{}, errors.New("etkinlik güncellenemedi")
	}
	return eventData, nil
}

func (s Service) DeleteEvent(eventID uint) error {
	err := s.repo.DeleteEvent(eventID)
	if err != nil {
		return errors.New("etkinlik silinemedi: " + err.Error())
	}
	return nil
}
