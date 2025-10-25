package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"luny.dev/sakila/w4-jwt/internal/middlewares"
	"luny.dev/sakila/w4-jwt/internal/models"
	"luny.dev/sakila/w4-jwt/internal/routes"
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

	EnsureEnvExists("JWT_SECRET")

	// Setup MySQL
	mysqlDsn := EnsureEnvExists("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("fatal: mysql database cannot be connected to")
	}

	// Do migrations
	err = db.AutoMigrate(&models.BlacklistedToken{}, &models.User{})
	if err != nil {
		log.Println("warn: migration failed. try resetting mysql database")
	}

	srv := gin.New()

	srv.Use(gin.Logger())

	{
		r := srv.Group("")

		jwtMiddleware := middlewares.CheckJWT(db)

		authHandler := routes.AuthHandler{DB: db}
		r.POST("/login", authHandler.Login)
		r.POST("/register", authHandler.Register)
		r.POST("/logout", jwtMiddleware, authHandler.Logout)

		profileHandler := routes.ProfileHandler{DB: db}
		r.GET("/profile", jwtMiddleware, profileHandler.GetProfile)
	}

	err = srv.Run(":80")
	if err != nil {
		log.Fatalln("fatal: failed to expose to port 80. taken?")
	}
}
