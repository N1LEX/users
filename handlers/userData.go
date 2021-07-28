package handlers

import (
	"butaforia.io/models"
	"butaforia.io/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func AuthUserData(c *gin.Context) {
	tokenData := c.GetHeader("Authorization")
	tokenValue := strings.Split(tokenData, " ")[1]
	tokenClaims := token.ParseTokenClaims(tokenValue)

	if !tokenClaims.IsValidAt(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authenticated",
		})
		return
	}

	u, err := models.GetByUsername(tokenClaims.Audience[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, u.GetShortData())
}
