package models

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Book{}, &Borrowing{})
}
