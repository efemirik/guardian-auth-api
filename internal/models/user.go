package models

import (
	"time"
	"gorm.io/gorm"
)

// User represents the users table in the database
type User struct {
	ID           uint           `gorm:"primaryKey"`
	Email        string         `gorm:"uniqueIndex;not null"`
	PasswordHash string         `gorm:"not null"` // Stored as securely hashed string (e.g., Argon2id)
	Role         string         `gorm:"default:'user'"` // Role-Based Access Control (RBAC) identifier
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}