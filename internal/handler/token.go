package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/services"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) {
	url := fmt.Sprintf("%s/access_token", config.GetEnv("reddit.url"))

	usename := "Piifia"
	password := "123456789asd*"
	client_secret := "1ptla1RhgWKykD9f53RaSyfVHj-FgA"
	client_id := "9vDiuDZcj1CNfWBn7hhUbw"

	authString := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))

	data := fmt.Sprintf("grant_type=password&username=%s&password=%s", usename, password)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"bar":          "Authorization " + authString,
	}

	services.Auth(url, headers, data, c)
}
