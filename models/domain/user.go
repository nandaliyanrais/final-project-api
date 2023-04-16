package domain

import (
	"time"

	"gorm.io/gorm"

	"mygram-api/helpers"
)

// User represents the model of a user
type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"not null;uniqueIndex"`
	Age       int        `gorm:"not null"`
	Email     string     `gorm:"not null;uniqueIndex"`
	Password  string     `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	
	user.Password = helpers.Hash(user.Password)

	return nil
}
