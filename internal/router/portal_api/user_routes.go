package router

import (
	"github.com/cbglabs/nexusgate-pbmapi/internal/db"
	handler "github.com/cbglabs/nexusgate-pbmapi/internal/handler/portal_api"
	service "github.com/cbglabs/nexusgate-pbmapi/internal/service/portal_api"

	"github.com/gin-gonic/gin"
)

// NewUserRouter sets up the routes for user-related operations.
// It defines the endpoints for adding, listing, retrieving, updating, and deleting users.
func NewUserRouter(routerGroup *gin.RouterGroup, authMiddleware gin.HandlerFunc) *gin.RouterGroup {
	dbClient, _ := db.NewDynamoDBClient()
	userService := service.NewUserService(dbClient)
	userHandler := handler.NewUserHandler(userService)

	// Create a sub-router for user-related routes.
	routes := routerGroup.Group("/users", authMiddleware)

	// Define the POST route for adding a new user.
	// Endpoint: POST /user/
	routes.POST("/", userHandler.CreateUser)

	// Define the GET route for retrieving a specific user by their "sk".
	// Endpoint: GET /user/:id
	routes.GET("/:id", userHandler.ReadUser)

	// Define the PUT route for updating a specific user by their "sk".
	// Endpoint: PUT /user/:id
	routes.PUT("/:id", userHandler.UpdateUser)

	// Define the DELETE route for deleting a specific user by their "sk".
	// Endpoint: DELETE /user/:id
	routes.DELETE("/:id", userHandler.DeleteUser)

	// Define the GET route for listing all users.
	// Endpoint: GET /user/list
	routes.GET("/list", userHandler.ListUsers)

	// Return the routes group instead of routerGroup
	return routes
}
