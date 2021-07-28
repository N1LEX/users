package main

import (
	"github.com/gin-gonic/gin"

	h "butaforia.io/handlers"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/token/new", h.NewToken)
	r.POST("/token/refresh", h.RefreshAccessToken)
	r.POST("/user/new", h.UserCreateHandler)
	r.POST("/token/verify", h.IsExpiredToken)
	r.GET("/user/profile", h.AuthUserData)
	return r
}
