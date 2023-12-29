package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/database/models/token"
	"RedditAutoPost/internal/request"
	"RedditAutoPost/internal/services"
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ObtainAccessToken(c *gin.Context, cfg *config.Config) (string, error) {
	token, _ := GetAccessToken(c, cfg)
	/*if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}*/
	fmt.Print(token)

	if !tokenIsValid(token) {
		return requestByToken(c, cfg)
	}

	return token.Token, nil
}

func secondsToDate(seconds float64) time.Time {
	return time.Now().Add(time.Duration(int64(seconds)) * time.Second)
}

func tokenIsValid(tok token.Tokens) bool {
	return time.Now().Unix() < tok.Expiration.Unix()
}

func requestByToken(c *gin.Context, cfg *config.Config) (string, error) {
	redditCredential, err := services.GetCredentials()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/v1/access_token", cfg.Reddit.Url)
	authString := base64.StdEncoding.EncodeToString([]byte(redditCredential.ClientID + ":" + redditCredential.ClientSecret))

	/*data := fmt.Sprintf("grant_type=password&username=%s&password=%s",
		redditCredential.Username,
		redditCredential.Password,
	)

	// data := []byte(`{"grant_type": "password", username , password }`)
	// data := []byte(`{"grant_type":"password","password":"` + redditCredential.Password + `","username":"` + redditCredential.Username + `"}`)
	// data := string([]byte(`{"grant_type":"password","password":"` + redditCredential.Password + `","username":"` + redditCredential.Username + `"}`))*/

	/*data := []byte(`{"grant_type:password, username":"` + redditCredential.Username + `","password":"` + redditCredential.Password + `"}`)
	bodyReader := bytes.NewReader(data)*/

	/*data, _ := json.Marshal(map[string]string{
		"grant_type": "password",
		"username":   redditCredential.Username,
		"password":   redditCredential.Password,
	})*/

	jsonData := []byte(`{"grant_type":"password", "username":"` + redditCredential.Username + `", "password":"` + redditCredential.Password + `"}`)

	responseBody := bytes.NewBuffer(jsonData)

	headers := map[string]string{
		// "Content-Type":  "application/x-www-form-urlencoded",
		"Content-Type":  "application/json; charset=utf-8",
		"Authorization": "Basic " + authString,
	}

	status, result := request.Post(url, headers, responseBody, c)

	c.JSON(status, result)

	if status == 200 {
		accessToken, ok := result["access_token"].(string)
		if !ok {
			log.Println("No se encontró el token de acceso")
			return "", nil
		}

		date := secondsToDate(result["expires_in"].(float64))

		tokenService, err := services.NewTokenService()
		if err != nil {
			return "", fmt.Errorf("error al crear el servicio de tokens: %v", err)
		}

		tokenService.CreateAccessToken(accessToken, date)

		return accessToken, nil
	}

	log.Printf("La solicitud no fue exitosa. Estado: %d\n", status)
	return "", nil
}

func GetAccessToken(c *gin.Context, cfg *config.Config) (token.Tokens, error) {
	tokenService, err := services.NewTokenService()
	if err != nil {
		return token.Tokens{}, fmt.Errorf("error al crear el servicio de tokens: %v", err)
	}

	tokens, err := tokenService.GetToken()
	if err != nil {
		return token.Tokens{}, fmt.Errorf("error al obtener el token: %v", err)
	}

	c.JSON(200, tokens)

	return tokens, nil
}

func DeleteToken(c *gin.Context, cfg *config.Config) error {
	tokenService, err := services.NewTokenService()
	if err != nil {
		return fmt.Errorf("error al crear el servicio de tokens: %v", err)
	}

	tokenService.DeleteToken()

	return nil
}
