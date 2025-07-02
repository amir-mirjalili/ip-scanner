package models

import (
	"time"
)

type ScanResult struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ScanID     uint      `gorm:"not null" json:"scanId"`
	AssetID    uint      `gorm:"not null" json:"assetId"`
	DetectedAt time.Time `gorm:"not null" json:"detectedAt"`

	Scan  Scan  `json:"scan,omitempty"`
	Asset Asset `json:"asset,omitempty"`
}
