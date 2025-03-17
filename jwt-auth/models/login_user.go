package models

import (
	"errors"
	"go/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) LoginCheck(db *gorm.DB, email, password string) (string, error) {
	u.Email = email
	if err := u.BeforeLogin(db, password); err != nil {
		return "", err
	}
	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", errors.New("error generating token occurs")
	}
	return token, nil
}
func (u *User) BeforeLogin(db *gorm.DB, password string) error {
	if err := db.Where("email - ?", u.Email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user is not found")
		}
		return err
	}
	if err := VerifyPassword(password, u.Password); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.New("password is incorrect")
		}
		return err
	}
	return nil
}
