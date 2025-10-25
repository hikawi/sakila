// Package middlewares provides middleware routing
package middlewares

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"luny.dev/sakila/w4-jwt/internal/models"
)

var bearerRe *regexp.Regexp = regexp.MustCompile("Bearer (.+)")

func CheckJWT(db *gorm.DB) func(*gin.Context) {
	return func(g *gin.Context) {
		authHeader := g.Request.Header.Get("Authorization")
		authSplit := bearerRe.FindStringSubmatch(authHeader)

		if len(authSplit) != 2 || !strings.HasPrefix(authHeader, "Bearer ") {
			g.AbortWithStatusJSON(401, gin.H{"message": "missing or invalid jwt"})
			return
		}

		token, err := jwt.Parse(authSplit[1], func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil || !token.Valid {
			g.AbortWithStatusJSON(401, gin.H{"message": "missing or invalid jwt"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"]
		jti := claims["jti"].(string)

		expiredAt, err := claims.GetExpirationTime()
		if err != nil {
			g.AbortWithStatusJSON(401, gin.H{"message": "missing or invalid jwt"})
			return
		}
		g.Set("expires_at", expiredAt)

		// Check against the DB for two cases:
		// Is the token blacklisted?
		// Is the ID or username different in the DB?
		ctx := context.Background()

		_, err = gorm.G[models.BlacklistedToken](db).Where("jti = ?", jti).First(ctx)
		if err == nil {
			g.AbortWithStatusJSON(401, gin.H{"message": "missing or invalid jwt"})
			return
		}

		user, err := gorm.G[models.User](db).Where("id = ?", id).First(ctx)
		if err != nil {
			g.AbortWithStatusJSON(401, gin.H{"message": "missing or invalid jwt"})
			return
		}

		g.Set("user", user)
		g.Set("jti", jti)
		g.Next()
	}
}
