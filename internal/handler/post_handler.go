package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/request"
	"bytes"
	"encoding/json"
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

	/*data := fmt.Sprintf("title=%s&text=%s&sr=%s&kind=%s",
		title,
		text,
		sr,
		kind,
	)*/

	data, _ := json.Marshal(map[string]string{
		"title": title,
		"text":  text,
		"sr":    sr,
		"kind":  kind,
	})
	responseBody := bytes.NewBuffer(data)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Bearer " + token.Token,
	}

	status, result := request.Post(url, headers, responseBody, c)

	c.JSON(status, result)

}
