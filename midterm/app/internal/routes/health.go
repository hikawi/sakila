// @BasePath /api/v1

// Package routes provides some routes I guess
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth returns a health
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func GetHealth(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"message": "healthy"})
}
