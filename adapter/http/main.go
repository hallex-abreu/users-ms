package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hallex-abreu/users-ms/adapter/http/actuator"
	"github.com/hallex-abreu/users-ms/adapter/http/authentication"
	"github.com/hallex-abreu/users-ms/adapter/http/users"
	"github.com/hallex-abreu/users-ms/database"
	"github.com/joho/godotenv"
)

func Init() {
	router := gin.Default()

	router.GET("/health", actuator.Health)
	router.GET("/users", users.Index)
	router.GET("/users/:id", users.Show)
	router.POST("/users", users.Store)
	router.POST("/login", authentication.Login)
	router.POST("/recover-password", authentication.RecoverPassword)
	router.POST("/reset-password", authentication.ResetPassword)

	database.Connection()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Run()
}
