package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status          string     `json:"status" gorm:"default:pending"`
	TotalPrice      float64    `json:"total_price"`
	ShippingAddress string     `json:"shipping_address"`
	City            string     `json:"city"`
	PostalCode      string     `json:"postal_code"`
	Province        string     `json:"province"`
	UserID          uint       `json:"user_id"`
	User            User       `json:"user" gorm:"foreignKey:UserID"`
	Items           []CartItem `json:"items" gorm:"foreignKey:OrderID"`
}
