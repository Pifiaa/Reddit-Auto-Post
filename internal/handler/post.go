package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	url := fmt.Sprintf("%s/api/submit", config.GetEnv("reddit.oauth"))

	services.Auth(url, c)
}
