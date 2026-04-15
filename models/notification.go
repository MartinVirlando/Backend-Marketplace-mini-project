package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Title       string         `json:"title"`
	Message     string         `json:"message"`
	Type        string         `json:"type"`
	IsRead      bool           `json:"is_read"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	ReferenceID *uint          `json:"reference_id"`
}
