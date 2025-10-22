package models

import "time"

type Language struct {
	LanguageID uint8     `gorm:"column:language_id;primaryKey;not null;autoIncrement" json:"language_id"`
	Name       string    `gorm:"column:name;size:20;not null" json:"name"`
	LastUpdate time.Time `gorm:"column:last_update;not null;autoUpdateTime" json:"last_updated"`
}

func (Language) TableName() string {
	return "language"
}
