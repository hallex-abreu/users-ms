package authentication

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hallex-abreu/users-ms/adapter/http/authentication/dtos"
	"github.com/hallex-abreu/users-ms/adapter/mail"
	"github.com/hallex-abreu/users-ms/database"
	"github.com/hallex-abreu/users-ms/entities"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body dtos.UserDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entities.Users
	result := database.DB.Where("email = ?", body.Email).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No user with this email."})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email or password incorrect."})
		return
	}

	token, err := CreateToken(int16(user.ID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	c.JSON(http.StatusOK, token)
}

func RecoverPassword(c *gin.Context) {
	var body dtos.UserDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entities.Users
	result := database.DB.Where("email = ?", body.Email).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No user with this email."})
		return
	}

	token, err := bcrypt.GenerateFromPassword([]byte(user.Email), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error generate token."})
		return
	}

	user.PasswordResetToken = string(token)
	user.PasswordResetExpires = time.Now().Add(time.Hour * 1).Unix()

	database.DB.Save(&user)

	mail.Send(user.Name, user.Email, string(token))

	c.JSON(http.StatusOK, gin.H{"message": "Enviado com sucesso! Verifique sua caixa de email."})
}

func CreateToken(userid int16) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
