// Package middlewares provides a set of middlewares to be directly used on GIN routing.
package middlewares

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"luny.dev/sakila/w4/internal/models"
)

func APIKeyAuthorization(db *gorm.DB) func(*gin.Context) {
	return func(g *gin.Context) {
		headers := g.Request.Header
		apiKeys := headers.Values("x-api-key")

		if len(apiKeys) != 1 {
			g.AbortWithStatusJSON(401, gin.H{"message": "missing or invalid x-api-key header"})
			return
		}

		_, err := db.DB()
		if err != nil {
			g.AbortWithStatusJSON(500, gin.H{"message": "unable to connect to db"})
			return
		}

		hasher := sha256.New()
		hasher.Write([]byte(apiKeys[0]))
		hashedKey := hex.EncodeToString(hasher.Sum(nil))

		ctx := context.Background()
		user, err := gorm.G[models.User](db).Where("token = ?", hashedKey).First(ctx)
		if err != nil {
			g.AbortWithStatusJSON(401, gin.H{"message": "missing or invalid x-api-key header"})
			return
		}

		g.Keys = make(map[any]any)
		g.Keys["user"] = user
	}
}
