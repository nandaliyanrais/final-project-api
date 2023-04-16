package domain

import "time"

// Photo represents the model of a photo
type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	UpdatedAt time.Time
	CreatedAt time.Time
}