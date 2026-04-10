package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = StringArray{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("failed to scan StringArray")
	}

	if len(bytes) == 0 {
		*s = StringArray{}
		return nil
	}

	return json.Unmarshal(bytes, s)
}

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Stock       int            `json:"stock"`

	Status string `json:"status" gorm:"default:pending"`

	Images StringArray `json:"images" gorm:"type:text;default:'[]'"`

	SellerID uint `json:"seller_id"`
	Seller   User `json:"seller" gorm:"foreignKey:SellerID"`

	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
}
