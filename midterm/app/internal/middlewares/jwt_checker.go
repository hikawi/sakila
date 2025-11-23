// Package middlewares provides a set of ready to use middlewares to bind to GIN.
package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"luny.dev/sakila/midterms/internal/utils"
)

// CheckJWT checks for the access token in the request. If it is invalid,
// it stops the request and returns 401.
func CheckJWT(g *gin.Context) {
	authHeader := g.GetHeader("Authorization")
	if authHeader == "" {
		g.AbortWithStatus(401)
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		g.AbortWithStatus(401)
		return
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(utils.FatalEnv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		g.AbortWithStatus(401)
		return
	}

	g.Next()
}
