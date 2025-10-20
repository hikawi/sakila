package models

import "time"

type Language struct {
	LanguageID uint16    `gorm:"primaryKey;not null;autoIncrement" json:"language_id"`
	Name       string    `gorm:"size:20;not null" json:"name"`
	LastUpdate time.Time `gorm:"not null;autoUpdateTime"`
}

func (Language) TableName() string {
	return "language"
}
