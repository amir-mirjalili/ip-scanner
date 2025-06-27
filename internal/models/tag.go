package models

import "time"

type Tag struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"size:100;unique;not null"`
	Assets    []Asset `gorm:"many2many:asset_tags"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
