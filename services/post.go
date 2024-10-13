package services

import (
	"github.com/onurravli/go-blog/db"
	"github.com/onurravli/go-blog/models"
)

func GetPosts() []models.Post {
	posts := []models.Post{}
	db.DB.Find(&posts)
	return posts
}

func GetPost(id int) models.Post {
	post := models.Post{}
	db.DB.First(&post, id)
	return post
}

func CreatePost(post models.Post) models.Post {
	if post.Title == "" || post.Body == "" {
		panic("Title and Body are required")
	}
	db.DB.Create(&post)
	return post
}
