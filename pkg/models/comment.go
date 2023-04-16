package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Author string
	Title  string `json:"title"`
	Book   string
}
