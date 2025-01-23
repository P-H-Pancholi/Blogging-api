package handlers

import (
	"net/http"
	"strconv"

	"github.com/P-H-Pancholi/Blogging-api/pkg/models"
	"github.com/P-H-Pancholi/Blogging-api/pkg/services"
	"github.com/gin-gonic/gin"
)

func GetAllPostsHandler(c *gin.Context) {
	posts, code, err := services.GetAllPosts()
	if err != nil {
		c.AbortWithError(code, err)
	}
	c.IndentedJSON(http.StatusOK, posts)
}

func CreatePostHandler(c *gin.Context) {
	var post *models.Post
	if err := c.BindJSON(&post); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if code, err := services.CreatePost(post); err != nil {
		c.AbortWithError(code, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, post)
}

func UpdatePostHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	var post *models.Post
	if err := c.BindJSON(&post); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if code, err := services.UpdatePost(post, uint(id)); err != nil {
		c.AbortWithError(code, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, post)
}
