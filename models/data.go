package models

import "time"

// Database type containing data

type Hashes struct {
	ID        uint `gorm:"primaryKey"`
	Hash      string
	CreatedAt time.Time
}
