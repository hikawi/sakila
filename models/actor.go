package models

import (
	"time"
)

// Actor represents the 'actor' table in the database.
type Actor struct {
	// gorm.Model is typically used, but to exactly match the schema,
	// we will define the fields explicitly.

	// actor_id smallint(5) unsigned NOT NULL AUTO_INCREMENT
	ActorID uint16 `gorm:"primaryKey;autoIncrement;column:actor_id" json:"actor_id"`

	// first_name varchar(45) NOT NULL
	FirstName string `gorm:"column:first_name;size:45;not null" json:"first_name"`

	// last_name varchar(45) NOT NULL, KEY idx_actor_last_name
	LastName string `gorm:"column:last_name;size:45;not null;index:idx_actor_last_name" json:"last_name"`

	// last_update timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	LastUpdate time.Time `gorm:"column:last_update;not null;autoUpdateTime" json:"last_updated_at"`
}

// TableName overrides the table name used by Actor to 'actor'
func (Actor) TableName() string {
	return "actor"
}
