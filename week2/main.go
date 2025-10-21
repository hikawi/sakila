package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "luny.dev/sakila/w2/docs"
	"luny.dev/sakila/w2/internal/routes"
)

func EnsureEnvExists(path string) string {
	val, exists := os.LookupEnv(path)
	if !exists {
		log.Fatalf("fatal: can't find required env variable %s\n", path)
	}
	return val
}

// @title           Sakila API Week 2
// @version         1.0
// @description     The Go/Gin backend API for the Sakila database for week 2.
// @host            localhost:3001
// @BasePath        /
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
		log.Fatalln("fatal: mysql database cannot be connected to")
	}

	app := gin.Default()

	// API group.
	{
		r := app.Group("")

		films := routes.FilmsHandler{DB: db}
		r.GET("/films", films.GetFilms)
		r.GET("/films/:id", films.GetFilmID)
		r.POST("/films", films.PostFilm)
		r.PATCH("/films/:id", films.PatchFilm)
		r.DELETE("/films/:id", films.DeleteFilm)
	}

	// Root group.
	// Lowest place to prevent any overrides, such as public serves.
	{
		r := app.Group("")
		r.GET("/docs", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
		})
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	app.Run("0.0.0.0:80")
}
