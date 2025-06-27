package models

import "time"

type ScanResult struct {
	ID         uint      `gorm:"primaryKey"`
	ScanID     uint      `gorm:"not null"`
	AssetID    uint      `gorm:"not null"`
	DetectedAt time.Time `gorm:"not null"`

	Scan  Scan
	Asset Asset
}
