// Package routes provides the necessary routing for apitoken version of week 4.
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"luny.dev/sakila/w4/internal/models"
)

type ProfileRouter struct {
	DB *gorm.DB
}

func (r ProfileRouter) GetProfile(g *gin.Context) {
	_, err := r.DB.DB()
	if err != nil {
		g.AbortWithStatusJSON(500, gin.H{"message": "unable to connect to db"})
		return
	}

	user := g.Keys["user"].(models.User)

	// Drop out the Token field
	var response struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	}
	response.ID = user.ID
	response.Name = user.Name

	g.JSON(200, response)
}
