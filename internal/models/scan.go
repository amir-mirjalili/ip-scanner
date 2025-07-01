package models

import (
	"time"
)

type Scan struct {
	ID         uint      `gorm:"primaryKey"`
	CIDR       string    `gorm:"size:100;not null"`
	StartedAt  time.Time `gorm:"not null"`
	FinishedAt *time.Time
	Status     ScanStatus `gorm:"size:50;not null"`
	Results    []ScanResult
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ScanStatus string

const (
	ScanStatusInProgress ScanStatus = "InProgress"
	ScanStatusCompleted  ScanStatus = "Completed"
	ScanStatusFailed     ScanStatus = "Failed"
)
