package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
	Author      string
	Rating      float32
	Price       float32
	Language    string
	Image       string
	Category    string
}
