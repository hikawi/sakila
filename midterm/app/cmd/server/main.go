package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	docs "luny.dev/sakila/midterms/docs"
	"luny.dev/sakila/midterms/internal/routes"
	"luny.dev/sakila/midterms/internal/utils"
)

func main() {
	server := gin.Default()
	utils.InitLogger()
	utils.InitRabbitMQ("midterms")

	docs.SwaggerInfo.Description = "Hey"
	docs.SwaggerInfo.BasePath = "/v1"

	db, err := gorm.Open(postgres.Open(utils.FatalEnv("POSTGRES_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalln("fatal: can't connect to postgresql")
	}
	db.DB()

	{
		g := server.Group("/v1")

		g.GET("/health", routes.GetHealth)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.Run(":80")
}
