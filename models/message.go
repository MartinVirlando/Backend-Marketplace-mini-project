package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
	Message    string         `json:"message"`
	IsRead     bool           `json:"is_read"`
	SenderID   uint           `json:"sender_id"`
	Sender     User           `json:"sender" gorm:"foreignKey:SenderID"`
	ReceiverID uint           `json:"receiver_id"`
	Receiver   User           `json:"receiver" gorm:"foreignKey:ReceiverID"`
	ProductID  *uint          `json:"product_id"`
	Product    *Product       `json:"product" gorm:"foreignKey:ProductID"`
}
