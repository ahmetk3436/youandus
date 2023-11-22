package repository

import (
	"errors"
	"youandus/pkg/event/model"

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

func (r Repository) UpdateEvent(eventID uint, eventData model.Event) (model.Event, error) {
	eventData.ID = eventID

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&model.Event{}, "id = ?", eventID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			if err := tx.Create(&eventData).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&model.Event{}).Where("id = ?", eventID).Updates(&eventData).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return model.Event{}, err
	}

	var newEvent model.Event
	if err := r.db.First(&newEvent, "id = ?", eventID).Error; err != nil {
		return model.Event{}, err
	}

	return newEvent, nil
}

func (r Repository) GetEvent(eventID uint) (model.Event, error) {
	var eventData model.Event
	if err := r.db.Table("events").Where("id = ?", eventID).Find(&eventData).Error; err != nil {
		return model.Event{}, err
	}
	return eventData, nil
}

func (r Repository) DeleteEvent(eventID uint) error {
	if err := r.db.Where("id = ?", eventID).Delete(&model.Event{}).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) CreateEvent(event *model.Event) error {
	if err := r.db.Create(&event).Error; err != nil {
		return err
	}
	return nil
}
