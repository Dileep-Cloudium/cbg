package router

import "github.com/gin-gonic/gin"

// Health check route
// @Summary Health check
// @Description Check the health of the service.
// @Tags Default
// @Accept json
// @Produce json
// @Success 200 {object} object{status=string} "Health check response"
// @Router /api/v1/health [get]
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// NewHealthRouter sets up the routes for health APIs
func NewHealthRouter(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	// Create a sub-router for health routes.
	routes := routerGroup.Group("/health")

	// Define the GET route for health check
	// Endpoint: GET /health
	routes.GET("", healthCheck)

	// Return the updated router group.
	return routes
}
