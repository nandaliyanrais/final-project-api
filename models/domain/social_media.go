package domain

import "time"

// SocialMedia represents the model of a social media
type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	SocialMediaUrl string `gorm:"not null;type:text"`
	UserID         uint   `gorm:"not null"`
	UpdatedAt      time.Time
	CreatedAt      time.Time
	User           User `gorm:"foreignKey:UserID"`
}