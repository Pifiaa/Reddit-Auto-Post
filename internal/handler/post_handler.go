package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context, cfg *config.Config) {
	token, _ := GetAccessToken(c, cfg)

	title := "Titulo de prueba"
	text := "Mensaje de prueba"
	sr := "pifiar4"
	kind := "self"

	url := fmt.Sprintf("%s/submit", cfg.Reddit.Oauth)

	data := fmt.Sprintf("title=%s&text=%s&sr=%s&kind=%s",
		title,
		text,
		sr,
		kind,
	)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Bearer " + token,
	}

	status, result := request.Post(url, headers, data, c)

	c.JSON(status, result)
}
