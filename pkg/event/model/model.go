package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventName        string `json:"eventName" gorm:"not null"`
	EventDescription string `json:"eventDescription" gorm:"not null"`
	EventDate        string `json:"eventDate" gorm:"not null"`
	EventLocation    string `json:"eventLocation" gorm:"not null"`
	EventType        string `json:"eventType" gorm:"not null"`
	Organizer        string `json:"organizer" gorm:"not null"`
	ContactEmail     string `json:"contactEmail" gorm:"not null"`
	ContactPhone     string `json:"contactPhone" gorm:"not null"`
	Website          string `json:"website" gorm:"not null"`
	Capacity         int    `json:"capacity" gorm:"not null"`
	UserID           uint   `json:"user_id" gorm:"not null"`
}

func (e Event) GetEventName() string {
	return e.EventName
}

func (e Event) GetEventDescription() string {
	return e.EventDescription
}

func (e Event) GetEventDate() string {
	return e.EventDate
}

func (e Event) GetEventLocation() string {
	return e.EventLocation
}

func (e Event) GetEventType() string {
	return e.EventType
}

func (e Event) GetOrganizer() string {
	return e.Organizer
}

func (e Event) GetContactEmail() string {
	return e.ContactEmail
}

func (e Event) GetContactPhone() string {
	return e.ContactPhone
}

func (e Event) GetWebsite() string {
	return e.Website
}

func (e Event) GetCapacity() int {
	return e.Capacity
}
