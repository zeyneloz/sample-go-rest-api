package models

import (
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// User struct describes the user model for GORM.
type User struct {
	ID       int    `gorm:"primary_key"`
	Email    string `gorm:"type:varchar(120);unique_index"`
	Name     string `gorm:"size:70"`
	Password []byte `gorm:"size:128"`
}

// SetPassword of user with given raw password.
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.NewInternalError(err, "Generate password error")
	}
	u.Password = hashedPassword
	return nil
}

// CheckPassword returns true if given rawPassword matches the password of the user.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	return err == nil
}
