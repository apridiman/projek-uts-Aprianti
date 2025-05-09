package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ISBN             string `gorm:"uniqueIndex"`
	Title            string
	Author           string
	YearPublished    int
	QuantityAvailable int
}
