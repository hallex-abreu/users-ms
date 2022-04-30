package users

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hallex-abreu/users-ms/adapter/http/users/dtos"
	"github.com/hallex-abreu/users-ms/database"
	"github.com/hallex-abreu/users-ms/entities"
	"golang.org/x/crypto/bcrypt"
)

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	filter := c.DefaultQuery("filter", "")

	var users []entities.Users
	database.DB.Where("name LIKE ?", "%"+filter+"%").Or("email LIKE ?", "%"+filter+"%").Offset(limit * (page - 1)).Limit(limit).Find(&users)
	var count int64
	var usersCount []entities.Users
	database.DB.Where("name LIKE ?", "%"+filter+"%").Or("email LIKE ?", "%"+filter+"%").Find(&usersCount).Count(&count)
	var total_elements = math.Ceil(float64(count) / float64(limit))

	if total_elements == 0 {
		total_elements = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"content":         users,
		"page":            page,
		"limit":           limit,
		"total_elements":  count,
		"number_elements": len(users),
		"tota_pages":      total_elements,
	})
}

func Show(c *gin.Context) {

	id := c.Params.ByName("id")

	var user entities.Users
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
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

	c.JSON(http.StatusOK, gin.H{"content": user})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
