package seeder

import (
	"library-app/models"
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Create admin user
	admin := models.User{
		Username:    "admin",
		PasswordHash: "admin123@", // Use bcrypt hash in real application
		FullName:    "Admin User",
		Role:        "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("Could not seed admin user:", err)
	}

	// Add books and other necessary records...
}
