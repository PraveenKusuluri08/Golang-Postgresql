package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	Title string
	Body  string
}
