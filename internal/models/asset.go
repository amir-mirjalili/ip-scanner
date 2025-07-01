package models

import "time"

type Asset struct {
	ID         uint    `gorm:"primaryKey"`
	IPAddress  string  `gorm:"size:100;not null"`
	MACAddress *string `gorm:"size:100"`
	Hostname   *string `gorm:"size:255"`
	OS         *string `gorm:"size:255"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
