package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hallex-abreu/users-ms/adapter/http/actuator"
	"github.com/hallex-abreu/users-ms/adapter/http/users"
	"github.com/hallex-abreu/users-ms/database"
)

func Init() {
	router := gin.Default()

	router.GET("/health", actuator.Health)
	router.GET("/users", users.Index)
	router.POST("/users", users.Store)

	database.Connection()

	router.Run()
}
