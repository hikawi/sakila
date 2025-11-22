// Package routes provides some routes I guess
package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"luny.dev/sakila/midterms/internal/utils"
)

type GetHealthResponse struct {
	Message string `json:"message"`
}

// GetHealth returns a health
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} routes.GetHealthResponse
// @Router /health [get]
func GetHealth(g *gin.Context) {
	utils.LogJSON(gin.H{"timestamp": time.Now().UnixMilli(), "action": "check_health"})
	g.JSON(http.StatusOK, gin.H{"message": "healthy"})
}
