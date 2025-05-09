package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

func MigrateEnums(db *gorm.DB) error {
	// Membuat tipe enum untuk Role
	err := db.Exec("CREATE TYPE user_role AS ENUM ('admin', 'member')").Error
	if err != nil {
		return fmt.Errorf("failed to create user_role enum type: %w", err)
	}

	// Membuat tipe enum untuk Status
	err = db.Exec("CREATE TYPE user_status AS ENUM ('active', 'inactive')").Error
	if err != nil {
		return fmt.Errorf("failed to create user_status enum type: %w", err)
	}

	return nil
}

