package models

import (
	"time"

	"gorm.io/gorm"
)

type Borrowing struct {
	gorm.Model
	MemberID   uint
	BookID     uint
	BorrowDate time.Time
	ReturnDate *time.Time
	Status     string `gorm:"type:enum('borrowed','returned');default:'borrowed'"`
}
