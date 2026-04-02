package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"uniqueIndex" json:"name"`
	Icon string `json:"icon,omitempty"`
}
