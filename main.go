package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/onurravli/go-blog/controllers"
	"github.com/onurravli/go-blog/db"
	"github.com/onurravli/go-blog/models"
)

func main() {
	db.Init()
	posts := []models.Post{}
	db.DB.Find(&posts)
	fmt.Println(posts)

	r := gin.Default()
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.Run()
}
