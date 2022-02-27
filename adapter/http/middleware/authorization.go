package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token"})
		c.Abort()
		return
	}

	token := strings.Split(authHeader, " ")
	if len(token) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalid."})
		c.Abort()
		return
	}

	// c.Set("body", m)
	c.Next()
}
