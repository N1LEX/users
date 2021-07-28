package handlers

import (
	"butaforia.io/forms"
	m "butaforia.io/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserCreateHandler(c *gin.Context) {
	var form forms.UserCreateForm

	// Binding POST form with UserCreateForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	// Validating form
	if errs := form.Validate(); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors":  errs,
		})
		return
	}

	// Create user
	user, err := m.CreateUser(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusCreated, user.GetShortData())
}