package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"luny.dev/awad-w1/routes"
)

func EnsureEnvExists(path string) string {
	val, exists := os.LookupEnv(path)
	if !exists {
		log.Fatalf("fatal: can't find required env variable %s\n", path)
	}
	return val
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("fatal: error loading .env file")
	}

	// Setup MySQL
	mysqlDsn := EnsureEnvExists("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatalf("fatal: mysql database cannot be connected to")
	}

	app := gin.Default()

	{
		r := app.Group("/api")
		r.GET("/health", routes.GetHealthFunc)

		actors := routes.ActorsHandler{DB: db}
		r.GET("/actors", actors.GetActors)
		r.GET("/actors/:id", actors.GetActorID)
		r.POST("/actors", actors.PostActor)
		r.DELETE("/actors/:id", actors.DeleteActor)
		r.PATCH("/actors/:id", actors.PatchActor)
	}

	app.Run("0.0.0.0:80")
}
