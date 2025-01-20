package router

import (
	"github.com/cbglabs/nexusgate-pbmapi/internal/db"
	handler "github.com/cbglabs/nexusgate-pbmapi/internal/handler/guiltyspark_api"
	service "github.com/cbglabs/nexusgate-pbmapi/internal/service/guiltyspark_api"

	"github.com/gin-gonic/gin"
)

// NewMemberRouter sets up the routes for member operations.
// It defines the endpoints for adding, listing, retrieving, updating, and deleting members.
func NewMemberRouter(routerGroup *gin.RouterGroup, authMiddleware gin.HandlerFunc) *gin.RouterGroup {
	dbClient, _ := db.NewGraphQLClient()
	memberService := service.NewMemberService(dbClient)
	memberHandler := handler.NewMemberHandler(memberService)

	// NOTE: Best Practices for PHI are as follows:
	// Use POST/PUT methods with request bodies instead of GET
	// Ensure data is sent in the request body, which is:
	// - Not logged by default
	// - Not stored in browser history
	// - Not visible in URLs
	// - Encrypted in transit with HTTPS

	routes := routerGroup.Group("/members", authMiddleware)

	// Define the POST route for querying members.
	// Endpoint: POST /members/
	routes.POST("/", memberHandler.ListMembers)

	// Define the POST route for querying members.
	// Endpoint: POST /members/list
	routes.POST("/list", memberHandler.ListMembers)

	// Define the POST route for checking member eligibility.
	// Endpoint: POST /members/check-eligibility
	routes.POST("/check-eligibility", memberHandler.CheckEligibility)

	return routes
}
