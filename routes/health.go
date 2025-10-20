// Package routes include a set of all routes under root.
package routes

import "github.com/gin-gonic/gin"

func GetHealthFunc(c *gin.Context) {
	c.JSON(200, gin.H{"status": "healthy", "version": "1", "path": "/v1"})
}
