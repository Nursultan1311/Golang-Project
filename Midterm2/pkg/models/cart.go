package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	BookId      int8
	UserId      int8
	Quantity    int8
	CreatedDate string
}
