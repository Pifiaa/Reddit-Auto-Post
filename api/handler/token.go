package handler

import (
	"RedditAutoPost/config"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(g *gin.Context) {
	usename := "Piifia"
	password := "123456789asd*"
	client_secret := "1ptla1RhgWKykD9f53RaSyfVHj-FgA"
	client_id := "9vDiuDZcj1CNfWBn7hhUbw"

	url := fmt.Sprintf("%s/access_token", config.GetEnv("reddit.url"))

	authString := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))

	data := fmt.Sprintf("grant_type=password&username=%s&password=%s", usename, password)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+authString)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, result)
}
