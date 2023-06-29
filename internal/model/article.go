package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title string
	Body  string
}
