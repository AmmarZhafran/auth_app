package models

import "gorm.io/gorm"

type User struct {
	Email        string `gorm:"unique"`
	PasswordHash string
	OTP          string `gorm:"-"`
	IsVerified   bool
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.IsVerified = false
	return
}
