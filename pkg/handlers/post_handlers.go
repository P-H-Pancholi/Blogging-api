package handlers

import (
	"net/http"

	"github.com/P-H-Pancholi/Blogging-api/pkg/models"
	"github.com/P-H-Pancholi/Blogging-api/pkg/services"
	"github.com/gin-gonic/gin"
)

func GetAllPostsHandler(c *gin.Context) {
	posts, err := services.GetAllPosts()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, posts)
}

func CreatePostHandler(c *gin.Context) {
	var post *models.Post
	if err := c.BindJSON(&post); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := services.CreatePost(post); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, post)
}
