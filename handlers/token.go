package handlers

import (
	f "butaforia.io/forms"
	m "butaforia.io/models"
	"butaforia.io/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// Generate JWT Token's for request user
func NewToken(c *gin.Context) {
	var form f.UserLoginForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username and password fields are required",
		})
		return
	}

	user, _ := m.GetByUsername(form.Username)
	if !user.IsValidCredentials(form.Username, form.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong password or username",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  user.NewAccessToken(),
		"refreshToken": user.NewRefreshToken(),
	})

}

func RefreshAccessToken(c *gin.Context) {
	refreshToken := c.PostForm("refreshToken")
	refreshTokenClaims := token.ParseTokenClaims(refreshToken)

	if !refreshTokenClaims.IsValidAt(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authenticated",
		})
		return
	}

	username := refreshTokenClaims.Audience[0]
	user, err := m.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": user.NewAccessToken(),
	})
}

// Check token expired
func IsExpiredToken(c *gin.Context) {
	tokenData := c.PostForm("token")
	tokenValue := strings.Split(tokenData, " ")[1]
	claims := token.ParseTokenClaims(tokenValue)
	c.JSON(http.StatusOK, gin.H{
		"expired": !claims.IsValidExpiresAt(time.Now()),
	})
}
