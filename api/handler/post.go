package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Post creado exitosamente",
	})
}
