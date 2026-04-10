package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Quantity  int            `json:"quantity"`
	UserID    uint           `json:"user_id"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	ProductID uint           `json:"product_id"`
	Product   Product        `json:"product" gorm:"foreignKey:ProductID"`
	OrderID   *uint          `json:"order_id,omitempty"`
}
