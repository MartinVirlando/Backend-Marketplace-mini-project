package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan StringArray")
	}
	return json.Unmarshal(bytes, s)
}

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`

	Status string `json:"status" gorm:"default:pending"`

	Images StringArray `json:"images" gorm:"type:text"`

	SellerID uint `json:"seller_id"`
	Seller   User `json:"seller" gorm:"foreignKey:SellerID"`

	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
}
