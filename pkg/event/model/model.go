package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventName        string `json:"eventName" gorm:"default:'empty'"`
	EventDescription string `json:"eventDescription" gorm:"default:'empty'"`
	EventDate        string `json:"eventDate" gorm:"default:'empty'"`
	EventLocation    string `json:"eventLocation" gorm:"default:'empty'"`
	EventType        string `json:"eventType" gorm:"default:'empty'"`
	Organizer        string `json:"organizer" gorm:"default:'empty'"`
	ContactEmail     string `json:"contactEmail" gorm:"default:'empty'"`
	ContactPhone     string `json:"contactPhone" gorm:"default:'empty'"`
	Website          string `json:"website" gorm:"default:'empty'"`
	Capacity         int    `json:"capacity" gorm:"default:0"`
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
