package actuator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealtBody struct {
	Status string `json: "status"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}
