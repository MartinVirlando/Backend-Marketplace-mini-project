package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating    int     `json:"rating"`
	Comment   string  `json:"comment"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
