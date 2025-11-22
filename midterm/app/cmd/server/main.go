package main

import (
	"github.com/gin-gonic/gin"
	"luny.dev/sakila/midterms/internal/routes"
)

func main() {
	server := gin.Default()

	{
		g := server.Group("/v1")

		g.GET("/health", routes.GetHealth)
	}

	server.Run(":80")
}
