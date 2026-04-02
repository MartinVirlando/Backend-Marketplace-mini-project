package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Quantity  int     `json:"quantity"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	OrderID   *uint   `json:"order_id,omitempty"`
}
