package storage

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	eventModel "youandus/pkg/event/model"
	profileModel "youandus/pkg/profile/model"
	userModel "youandus/pkg/users/model/user"
)

var db *gorm.DB // database

func init() {
	dsn := "ahmet:SXLV;jFPxT34i%VOYlUX#A6rN^1a;y@tcp(mysql:3306)/youanduseventplanner?charset=utf8mb4&parseTime=True&loc=Local"

	var conn *gorm.DB
	var err error
	for {
		conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
		if err == nil {
			break
		}
		fmt.Printf("Bağlantı hatası: %s\n", err.Error())
		time.Sleep(30 * time.Second)
	}

	db = conn
	err = db.AutoMigrate(&userModel.UserRegister{}, &profileModel.User{}, &eventModel.Event{})
	if err != nil {
		panic(err)
	}
	println("DB Successfully migrated !")
}

func GetDB() *gorm.DB {
	return db
}
