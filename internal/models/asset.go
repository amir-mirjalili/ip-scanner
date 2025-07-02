package models

import "time"

type Asset struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	IPAddress  string    `gorm:"size:100;not null" json:"ipAddress"`
	MACAddress *string   `gorm:"size:100" json:"macAddress,omitempty"`
	Hostname   *string   `gorm:"size:255" json:"hostname,omitempty"`
	OS         *string   `gorm:"size:255" json:"os,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
