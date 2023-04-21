package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Type     string `json:"category"`
	Password []byte
}
