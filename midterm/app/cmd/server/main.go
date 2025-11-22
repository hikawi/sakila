package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	docs "luny.dev/sakila/midterms/docs"
	"luny.dev/sakila/midterms/internal/routes"
	"luny.dev/sakila/midterms/internal/utils"
)

func Fatalenv(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("fatal: env variable %s not found\n", key)
	}

	return val
}

func main() {
	server := gin.Default()
	utils.InitLogger()

	docs.SwaggerInfo.Description = "Hey"
	docs.SwaggerInfo.BasePath = "/v1"

	db, err := gorm.Open(postgres.Open(Fatalenv("POSTGRES_DSN")), &gorm.Config{})
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
