package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar,omitempty"`
	Phone    string `json:"phone,omitempty"` 
}
