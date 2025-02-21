package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahar3/api-gateway/pkg/client"
)

// CreateTripHandler handles POST /trip requests.
func CreateTripHandler(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Forward the request to the Trip Planning Service
	response, err := client.CreateTrip(requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response from backend"})
		return
	}

	c.JSON(http.StatusOK, result)
}
