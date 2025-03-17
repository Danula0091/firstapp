package repo

import (
	"auth/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})
	return db
}
