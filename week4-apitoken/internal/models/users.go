// Package models provides a set of mapped models with GORM.
package models

type User struct {
	ID    int32  `gorm:"column:id;primaryKey;autoIncrement"`
	Name  string `gorm:"column:name;not null"`
	Token string `gorm:"column:token;not null"`
}

func (User) TableName() string {
	return "users"
}
