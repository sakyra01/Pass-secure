package models

import "time"

type Hashes struct {
	ID        uint `gorm:"primaryKey"`
	Hash      string
	CreatedAt time.Time
}
