package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hallex-abreu/users-ms/adapter/http/users/dtos"
	"github.com/hallex-abreu/users-ms/database"
	"github.com/hallex-abreu/users-ms/entities"
	"golang.org/x/crypto/bcrypt"
)

func Index(c *gin.Context) {
	var users []entities.Users
	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Store(c *gin.Context) {
	var body dtos.UserDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.Users{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashPassword),
	}

	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
