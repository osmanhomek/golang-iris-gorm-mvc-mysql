package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `gorm:"size:15"`
	Author string `gorm:"size:15"`
}
