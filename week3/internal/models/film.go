// Package models provides a list of all usable models in GORM MySQL.
package models

import (
	"time"

	"luny.dev/sakila/w3/utils"
)

/*
CREATE TABLE `film` (
  `film_id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text,
  `release_year` year(4) DEFAULT NULL,
  `language_id` tinyint(3) unsigned NOT NULL,
  `original_language_id` tinyint(3) unsigned DEFAULT NULL,
  `rental_duration` tinyint(3) unsigned NOT NULL DEFAULT '3',
  `rental_rate` decimal(4,2) NOT NULL DEFAULT '4.99',
  `length` smallint(5) unsigned DEFAULT NULL,
  `replacement_cost` decimal(5,2) NOT NULL DEFAULT '19.99',
  `rating` enum('G','PG','PG-13','R','NC-17') DEFAULT 'G',
  `special_features` set('Trailers','Commentaries','Deleted Scenes','Behind the Scenes') DEFAULT NULL,
  `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`film_id`),
  KEY `idx_title` (`title`),
  KEY `idx_fk_language_id` (`language_id`),
  KEY `idx_fk_original_language_id` (`original_language_id`),
  CONSTRAINT `fk_film_language` FOREIGN KEY (`language_id`) REFERENCES `language` (`language_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_film_language_original` FOREIGN KEY (`original_language_id`) REFERENCES `language` (`language_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1002 DEFAULT CHARSET=utf8;
*/

type Film struct {
	FilmID      uint16 `gorm:"primaryKey;autoIncrement;column:film_id;not null" json:"film_id"`
	Title       string `gorm:"size:255;not null;column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	ReleaseYear *int16 `gorm:"type:YEAR(4);default:null;column:release_year" json:"release_year"`

	LanguageID uint8    `gorm:"column:language_id;constraint:OnUpdate:CASCADE" json:"language_id"`
	Language   Language `gorm:"foreignKey:LanguageID;references:LanguageID" json:"language"`

	OriginalLanguageID *uint8    `gorm:"column:original_language_id;default:null;constraint:OnUpdate:CASCADE" json:"original_language_id"`
	OriginalLanguage   *Language `gorm:"foreignKey:OriginalLanguageID;default:null;references:LanguageID" json:"original_language"`

	RentalDuration  uint8           `gorm:"column:rental_duration;not null;default:3" json:"rental_duration"`
	RentalRate      float32         `gorm:"column:rental_rate;not null;default:4.99" json:"rental_rate"`
	Length          *uint16         `gorm:"column:length;default:null" json:"length"`
	ReplacementCost float32         `gorm:"column:replacement_cost;default:19.99;not null" json:"replacement_cost"`
	Rating          string          `gorm:"column:rating;type:enum('G','PG','PG-13','R','NC-17');default:G" json:"rating"`
	SpecialFeatures utils.StringSet `gorm:"column:special_features;type:set('Trailers','Commentaries','Deleted Scenes','Behind the Scenes');default:null" json:"special_features"`
	LastUpdate      time.Time       `gorm:"column:last_update;autoUpdateTime;not null"`
}

func (Film) TableName() string {
	return "film"
}
