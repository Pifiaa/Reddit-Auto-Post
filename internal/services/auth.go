package services

import (
	"RedditAutoPost/internal/http"
	"RedditAutoPost/internal/models/credentials"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(url string, c *gin.Context) {
	var firstCredential credentials.Credentials = GetCredentials()

	username := firstCredential.Username
	password := firstCredential.Password
	client_secret := firstCredential.ClientSecret
	client_id := firstCredential.ClientID

	authString := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))

	data := fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + authString,
	}

	http.Post(url, headers, data, c)
}

func RefreshToken(url string, c *gin.Context) {
	var firstCredential credentials.Credentials = GetCredentials()

	client_secret := firstCredential.ClientSecret
	client_id := firstCredential.ClientID

	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IlNIQTI1NjpzS3dsMnlsV0VtMjVmcXhwTU40cWY4MXE2OWFFdWFyMnpLMUdhVGxjdWNZIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ1c2VyIiwiZXhwIjoxNzAyNzUyMzIyLjEzNTQ4OCwiaWF0IjoxNzAyNjY1OTIyLjEzNTQ4NywianRpIjoiRmg4ZktZejRINW4wTi15bmt1RjR0el9vMnZmUHR3IiwiY2lkIjoiOXZEaXVEWmNqMUNOZldCbjdoaFVidyIsImxpZCI6InQyX296aTd4aHZzYiIsImFpZCI6InQyX296aTd4aHZzYiIsImxjYSI6MTcwMTM2Mzg3NjcyMiwic2NwIjoiZUp5S1Z0SlNpZ1VFQUFEX193TnpBU2MiLCJmbG8iOjl9.Cm7YPebxYga7JfXNMXTCYeiZngQfp_g8jTOYvWNVF0gvN1mHQZGfH_zu1oQJ4D4DYm4G7uLFWgAMydgcOt7eIrNJspHX8wz2OcAaIvWrKcGGzXozjoPlLuDLmLb0ShANmpM1RUeaY8TuMAIeDfD6naXpVVgBtPbPrYCk3MGIotoPCFjfubHpr2qubE6XqBDjCphGF7PNYKOdvFxiwoIs3TVC5NqaBajl6GJjHsYObScY1st-5TVJS8Ze-eG-l7-Y6cS1IO3xqQT8TzUvWhTfa7dQ0GPYbP4-WLQtG3tybHCtN4CadYlpoNjnWRO9IER38VBs7Wcp6EUgthti8aI0WQ"

	data := fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", token)
	authString := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + authString,
	}

	http.Post(url, headers, data, c)
}
