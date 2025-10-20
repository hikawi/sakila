package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"luny.dev/sakila/routes"
	v1 "luny.dev/sakila/routes/v1"
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
		log.Println("warn: error loading .env file, hopefully it is injected in")
	}

	// Setup MySQL
	mysqlDsn := EnsureEnvExists("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatalf("fatal: mysql database cannot be connected to")
	}

	app := gin.Default()

	// API v1 group.
	{
		r := app.Group("/v1")

		actors := v1.ActorsHandler{DB: db}
		r.GET("/actors", actors.GetActors)
		r.GET("/actors/:id", actors.GetActorID)
		r.POST("/actors", actors.PostActor)
		r.DELETE("/actors/:id", actors.DeleteActor)
		r.PATCH("/actors/:id", actors.PatchActor)

		films := v1.FilmsHandler{DB: db}
		r.GET("/films", films.GetFilms)
		r.GET("/films/:id", films.GetFilmID)
		r.POST("/films", films.PostFilm)
		r.PATCH("/films/:id", films.PatchFilm)
		r.DELETE("/films/:id", films.DeleteFilm)
	}

	// Root group.
	{
		r := app.Group("")
		r.GET("/health", routes.GetHealthFunc)
	}

	app.Run("0.0.0.0:80")
}
