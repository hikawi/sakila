package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"luny.dev/sakila/w4-jwt/internal/models"
)

type ProfileHandler struct {
	DB *gorm.DB
}

func (h ProfileHandler) GetProfile(g *gin.Context) {
	user, ok := g.Get("user")
	if !ok {
		g.AbortWithStatusJSON(401, gin.H{"message": "invalid user"})
		return
	}

	userModel := user.(models.User)
	g.JSON(200, gin.H{"id": userModel.ID, "username": userModel.Username})
}
