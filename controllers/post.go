package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/onurravli/go-blog/services"
)

func GetPosts(c *gin.Context) {
	posts := services.GetPosts()
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post := services.GetPost(id)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}
