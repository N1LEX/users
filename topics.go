package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	db "../database"
	m "../models"
)


func FetchAll(c *gin.Context) {
	var results []map[string]interface{}

	db.DB.
		Select("title, content, votes, comments").
		Model(&m.Topics{}).
		Scan(&results)

	c.JSON(200, gin.H{
		"results": results,
	})
}

func NewTopic(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	t := &m.Topics{Title: title, Content: content}
	db.DB.Create(t)
	c.JSON(http.StatusCreated, t)
}