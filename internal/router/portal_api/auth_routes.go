package router

import (
	handler "github.com/cbglabs/nexusgate-pbmapi/internal/handler/portal_api"

	"github.com/gin-gonic/gin"
)

// NewAuthRouter sets up the routes for auth APIs
func NewAuthRouter(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	// Create a sub-router for auth routes.
	routes := routerGroup.Group("/auth")

	// Define the POST route for login
	// Endpoint: POST /auth/login
	routes.POST("/login", handler.Login)

	// Define the POST route for register
	// Endpoint: POST /auth/register
	routes.POST("/register", handler.Register)

	// Return the updated router group.
	return routes
}

// NewPublicRouter sets up the routes for public APIs
func NewPublicRouter(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	// Create a sub-router for public routes.
	routes := routerGroup.Group("/public")

	// Define the GET route for listing all states
	// Endpoint: GET /public/states
	routes.GET("/states", handler.GetStates)

	// Return the updated router group.
	return routes
}
