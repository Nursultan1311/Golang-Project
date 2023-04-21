package models

import "gorm.io/gorm"

type Filter struct {
	gorm.Model
	From string
	To   string
}
