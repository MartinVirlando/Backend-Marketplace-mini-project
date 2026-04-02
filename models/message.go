package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Message    string   `json:"message"`
	IsRead     bool     `json:"is_read"`
	SenderID   uint     `json:"sender_id"`
	Sender     User     `json:"sender" gorm:"foreignKey:SenderID"`
	ReceiverID uint     `json:"receiver_id"`
	Receiver   User     `json:"receiver" gorm:"foreignKey:ReceiverID"`
	ProductID  *uint    `json:"product_id"`
	Product    *Product `json:"product" gorm:"foreignKey:ProductID"`
}
