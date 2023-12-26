package request

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(url string, headers map[string]string, data string, c *gin.Context) (int, map[string]interface{}) {

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))

	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}

	return http.StatusOK, result
}
