package storage

import (
	"fmt"
	"time"
	"youandus/pkg/profile/model"
	"youandus/pkg/users/model/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // database

func init() {
	dsn := "ahmet:SXLV;jFPxT34i%VOYlUX#A6rN^1a;y@tcp(localhost:3306)/youanduseventplanner?charset=utf8mb4&parseTime=True&loc=Local"

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
	db.AutoMigrate(&user.UserRegister{}, &model.User{}) // Database migration
}

func GetDB() *gorm.DB {
	return db
}
