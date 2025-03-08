package entities

import "time"

type Profile struct {
	Firstname string    `gorm:"not null"`
	Lastname  string    `gorm:"not null"`
	Birthday  time.Time `gorm:"not null"`
	AccountID uint
}
