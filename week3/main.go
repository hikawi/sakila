package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"luny.dev/sakila/w3/internal/routes"
	"luny.dev/sakila/w3/utils"
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

	utils.InitLogger("/var/log/week3/app.log")

	// Setup MySQL
	mysqlDsn := EnsureEnvExists("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("fatal: mysql database cannot be connected to")
	}

	g := gin.New()
	gin.DisableConsoleColor()

	g.Use(gin.Recovery())

	{
		r := g.Group("")

		films := routes.FilmsHandler{DB: db}
		r.GET("/films", films.GetFilms)
		r.GET("/films/:id", films.GetFilmID)
		r.POST("/films", films.PostFilm)
		r.PATCH("/films/:id", films.PatchFilm)
		r.DELETE("/films/:id", films.DeleteFilm)
	}

	g.Run(":80")
}
