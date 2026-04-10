package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Rating    int            `json:"rating"`
	Comment   string         `json:"comment"`
	UserID    uint           `json:"user_id"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	ProductID uint           `json:"product_id"`
	Product   Product        `json:"product" gorm:"foreignKey:ProductID"`
}
