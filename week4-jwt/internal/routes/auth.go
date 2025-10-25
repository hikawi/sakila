// Package routes provides a list of routing options for Week 4 JWT.
package routes

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"luny.dev/sakila/w4-jwt/internal/models"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h *AuthHandler) Register(g *gin.Context) {
	var request struct {
		Username string `binding:"required,min=3,max=64,alphanum" json:"username"`
		Password string `binding:"required,min=8,max=64" json:"password"`
	}
	err := g.ShouldBindJSON(&request)
	if err != nil {
		g.AbortWithStatusJSON(400, gin.H{"message": "bad request", "error": err.Error()})
		return
	}

	ctx := context.Background()

	// Check for existing username
	_, err = gorm.G[models.User](h.DB).
		Where("lower(username) = ?", strings.ToLower(request.Username)).
		First(ctx)
	if err == nil {
		g.AbortWithStatusJSON(409, gin.H{"message": "username exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		g.AbortWithStatusJSON(500, gin.H{"message": "failed to hash password"})
		return
	}

	// Create the user.
	user := models.User{
		Username: request.Username,
		Password: string(hashedPassword),
	}
	err = gorm.G[models.User](h.DB).Create(ctx, &user)
	if err != nil {
		g.AbortWithStatusJSON(500, gin.H{"message": "unable to create a user"})
		return
	}

	g.JSON(201, gin.H{"id": user.ID, "username": user.Username})
}

func (h *AuthHandler) Login(g *gin.Context) {
	var request struct {
		Username string `binding:"required,min=3,max=64,alphanum" json:"username"`
		Password string `binding:"required,min=8,max=64" json:"password"`
	}
	err := g.ShouldBindJSON(&request)
	if err != nil {
		g.AbortWithStatusJSON(400, gin.H{"message": "bad request", "error": err.Error()})
		return
	}

	ctx := context.Background()

	// Check for existing username
	user, err := gorm.G[models.User](h.DB).
		Where("lower(username) = ?", strings.ToLower(request.Username)).
		First(ctx)
	if err != nil {
		g.AbortWithStatusJSON(404, gin.H{"message": "username does not exist"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		g.AbortWithStatusJSON(403, gin.H{"message": "wrong password"})
		return
	}

	// Create the JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
		"jti":      uuid.New(),
	})
	jwt, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		g.AbortWithStatusJSON(500, gin.H{"message": "couldn't sign jwt", "error": err.Error()})
		return
	}

	g.JSON(200, gin.H{"token": jwt})
}

func (h *AuthHandler) Logout(g *gin.Context) {
	jti, ok := g.Get("jti")
	if !ok {
		g.AbortWithStatusJSON(204, gin.H{"message": "not logged in"})
		return
	}

	expiresAt, ok := g.Get("expires_at")
	if !ok {
		g.AbortWithStatusJSON(204, gin.H{"message": "invalid jwt"})
		return
	}

	ctx := context.Background()
	token := models.BlacklistedToken{
		JTI:       jti.(string),
		ExpiresAt: time.Unix(expiresAt.(*jwt.NumericDate).Unix(), 0),
	}
	err := gorm.G[models.BlacklistedToken](h.DB).Create(ctx, &token)
	if err != nil {
		g.AbortWithStatusJSON(204, gin.H{"message": "unable to blacklist token"})
		return
	}

	g.JSON(200, gin.H{"message": "blacklisted token"})
}
