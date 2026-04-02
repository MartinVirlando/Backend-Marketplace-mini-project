package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Title   string `json:"title"`
	Message string `json:"message"`
	Type    string `json:"type"`
	IsRead  bool   `json:"is_read"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`
}
