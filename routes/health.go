// Package routes include a set of all routes under /api
package routes

import "github.com/gin-gonic/gin"

func GetHealthFunc(c *gin.Context) {
	c.JSON(200, gin.H{"status": "healthy"})
}
