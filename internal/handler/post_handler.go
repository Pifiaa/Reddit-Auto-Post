package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context, cfg *config.Config) {
	token, _ := GetAccessToken(c, cfg)

	title := "Pinche reddit no tiene un sandbox para hacer pruebas"
	text := "La api de prueba de reddit no tiene un ambiente de pruebas"
	sr := "testeandoconpifia"
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

	c.JSON(200, url)

	status, result := request.Post(url, headers, data, c)

	c.JSON(status, result)
}
