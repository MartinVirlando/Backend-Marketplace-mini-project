package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name      string         `json:"name"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"-"`
	Role      string         `json:"role"`
	Avatar    string         `json:"avatar,omitempty"`
	Phone     string         `json:"phone,omitempty"`
}
