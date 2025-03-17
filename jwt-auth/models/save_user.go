package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	if err := u.BeforeSave(db); err != nil {
		return nil, err
	}
	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
func (u *User) BeforeSave(db *gorm.DB) error {
	var existedUser User
	if err := db.Where("email - ?", u.Email).First(&existedUser).Error; err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}
