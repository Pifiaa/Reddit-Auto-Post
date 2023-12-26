package handler

import (
	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) {
	/*// Accede a la variable del servidor (ginServer) desde el contexto
	server := c.MustGet("server").(*ginServer)

	// Accede a la base de datos desde la instancia de ginServer
	db := server.db

	// Resto del código..*/

	/*redditCredential, err := services.GetCredentials()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if len(redditCredential) == 0 {
		err = fmt.Errorf("No hay credenciales de acceso disponibles")
		c.JSON(http.StatusNotFound, gin.H{"Error: ": err.Error()})
		return
	}

	username := redditCredential[0].Username
	password := redditCredential[0].Password
	client_secret := redditCredential[0].ClientSecret
	client_id := redditCredential[0].ClientID
	authString := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))

	data := fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + authString,
	}

	url := fmt.Sprintf("%s/access_token", config.GetEnv("reddit.url"))

	status, result := request.Post(url, headers, data, c)

	if status == 200 {

		timestamp := result["expires_in"].(float64)
		myDate := time.Unix(timestamp, 0)
		services.CreateAccessToken(result["access_token"].(string), myDate)
	}*/
}
