package models

import "time"

type Asset struct {
	ID         uint    `gorm:"primaryKey"`
	IPAddress  string  `gorm:"size:100;not null"`
	MACAddress *string `gorm:"size:100"`
	Hostname   *string `gorm:"size:255"`
	OS         *string `gorm:"size:255"`
	Tags       []Tag   `gorm:"many2many:asset_tags"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
