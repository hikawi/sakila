package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "luny.dev/sakila/docs"
	"luny.dev/sakila/routes"
	v1 "luny.dev/sakila/routes/v1"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func EnsureEnvExists(path string) string {
	val, exists := os.LookupEnv(path)
	if !exists {
		log.Fatalf("fatal: can't find required env variable %s\n", path)
	}
	return val
}

// @title           Sakila API
// @version         1.0
// @description     The Go/Gin backend API for the Sakila database.
// @host            api.sakila.luny.dev
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
	// Lowest place to prevent any overrides, such as public serves.
	{
		r := app.Group("")
		r.GET("/health", routes.GetHealth)
		r.GET("/docs", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
		})
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	app.Run("0.0.0.0:80")
}
