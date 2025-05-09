package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique"`
	FullName   string
	PasswordHash string
	Role       UserRole `gorm:"type:user_role;default:'member'"` // Menggunakan enum
	Status     UserStatus `gorm:"type:user_status;default:'active'"` // Menggunakan enum
	CreatedAt  string
	UpdatedAt  string
}

type UserRole string
type UserStatus string

// Enum untuk Role
const (
	Admin  UserRole = "admin"
	Member UserRole = "member"
)

// Enum untuk Status
const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
)
