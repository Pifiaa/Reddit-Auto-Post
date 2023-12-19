package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) {
	url := fmt.Sprintf("%s/access_token", config.GetEnv("reddit.url"))

	services.Auth(url, c)
}
