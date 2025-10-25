// Package models provides a set of GORM-mapped models in MySQL
package models

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type BlacklistedToken struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement"`
	JTI       string    `gorm:"column:jti;uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"column:expires_at;index;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null"`
}

func PruneExpiredTokens(ctx context.Context, db *gorm.DB) {
	count, err := gorm.G[BlacklistedToken](db).Where("expired_at < ?", time.Now()).Delete(ctx)
	if err != nil {
		log.Println("warn: failed to prune expired tokens")
	} else {
		log.Printf("info: pruned %d tokens", count)
	}
}
