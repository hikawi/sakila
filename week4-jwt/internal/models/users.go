package models

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `gorm:"column:username;type:varchar(255) COLLATE utf8mb4_unicode_ci;uniqueIndex;not null"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime"`
}
