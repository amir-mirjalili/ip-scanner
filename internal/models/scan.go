package models

import (
	"time"
)

type Scan struct {
	ID         uint         `gorm:"primaryKey" json:"id"`
	CIDR       string       `gorm:"size:100;not null" json:"cidr"`
	StartedAt  time.Time    `gorm:"not null" json:"startedAt"`
	FinishedAt *time.Time   `json:"finishedAt,omitempty"`
	Status     ScanStatus   `gorm:"size:50;not null" json:"status"`
	Results    []ScanResult `json:"results,omitempty"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
}

type ScanStatus string

const (
	ScanStatusInProgress ScanStatus = "InProgress"
	ScanStatusCompleted  ScanStatus = "Completed"
	ScanStatusFailed     ScanStatus = "Failed"
)
