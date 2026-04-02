package config

import (
	"fmt"
	"log"

	"backend/models"
)

func MigrateDB() {
	fmt.Println("Migrating DB...")

	err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
		&models.Review{},
		&models.Message{},
		&models.Notification{},
	)

	if err != nil {
		log.Fatal("Failed to migrate DB", err)
	}

	fmt.Println("DB Migrated!")
}
