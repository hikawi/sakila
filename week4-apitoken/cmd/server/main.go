package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"luny.dev/sakila/w4/internal/middlewares"
	"luny.dev/sakila/w4/internal/routes"
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
		log.Fatalln("fatal: mysql database cannot be connected to")
	}

	srv := gin.New()

	srv.Use(gin.Logger())

	profilesRouter := routes.ProfileRouter{DB: db}
	srv.GET("/profile", middlewares.APIKeyAuthorization(db), profilesRouter.GetProfile)

	authRouter := routes.AuthRouter{DB: db}
	srv.POST("/register", authRouter.Register)

	err = srv.Run(":80")
	if err != nil {
		log.Fatalln("fatal: failed to expose to port 80. taken?")
	}
}
