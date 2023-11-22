package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName         string `json:"firstName" gorm:"default:'empty'"`
	SurName           string `json:"surName" gorm:"default:'empty'"`
	UserName          string `json:"userName" gorm:"default:'empty'"`
	UserEmail         string `json:"userEmail" gorm:"default:'empty'"`
	FoundYear         string `json:"foundYear" gorm:"default:'empty'"`
	OrganizationType  string `json:"organizationType" gorm:"default:'empty'"`
	Capacity          string `json:"capacity" gorm:"default:'empty'"`
	Adress            string `json:"address" gorm:"default:'empty'"`
	State             string `json:"state" gorm:"default:'empty'"`
	City              string `json:"city" gorm:"default:'empty'"`
	ZipCode           string `json:"zipCode" gorm:"default:'empty'"`
	AboutOrganization string `json:"aboutOrganization" gorm:"default:'empty'"`
	PhoneNumber       string `json:"phoneNumber" gorm:"default:'empty'"`
	WhatsappNumber    string `json:"whatsappNumber" gorm:"default:'empty'"`
	KipAdress         string `json:"kipAdress" gorm:"default:'empty'"`
	Facebook          string `json:"facebook" gorm:"default:'empty'"`
	Instagram         string `json:"instagram" gorm:"default:'empty'"`
	Twitter           string `json:"twitter" gorm:"default:'empty'"`
	LinkedIn          string `json:"linkedin" gorm:"default:'empty'"`
	UserID            uint   `json:"user_id" validate:"required" gorm:"unique"`
}

func (u User) GetUserFullName() string {
	return u.FirstName + " " + u.SurName
}
func (u User) GetUserEmail() string {
	return u.UserEmail
}
func (u User) GetUserName() string {
	return u.UserName
}
func (u User) GetLocation() interface{} {
	location := struct {
		state   string
		city    string
		zipCode string
	}{
		state:   u.State,
		city:    u.City,
		zipCode: u.ZipCode,
	}
	return location
}
func (u User) GetOrganization() interface{} {
	organization := struct {
		organizationType string
		capacity         string
		foundYear        string
	}{
		organizationType: u.OrganizationType,
		capacity:         u.Capacity,
		foundYear:        u.FoundYear,
	}
	return organization
}
func (u User) GetUserNameAndEmail() interface{} {
	nameEmail := struct {
		userName string
		email    string
	}{
		userName: u.UserName,
		email:    u.UserEmail,
	}
	return nameEmail
}
