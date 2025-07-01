package db

import (
	"fmt"
	"github.com/amir-mirjalili/ip-scanner/internal/models"
)

func AutoMigrate(database *Database) error {
	err := database.DB.AutoMigrate(
		&models.Asset{},
		&models.Scan{},
		&models.ScanResult{},
	)
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	fmt.Println("✅ Database migration completed successfully")
	return nil
}
