package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shahar3/api-gateway/handlers"
)

const (
	APIVersion = "v1"
)

// SetupRoutes defines all the API endpoints.
func SetupRoutes(router *gin.Engine) {
	// API versioning and grouping can be added here
	api := router.Group(fmt.Sprintf("/api/%s", APIVersion))
	{
		api.POST("/trip", handlers.CreateTripHandler)
	}
}
