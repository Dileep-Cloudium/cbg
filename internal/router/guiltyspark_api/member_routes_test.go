package router

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewMemberRouter(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	baseGroup := router.Group("/api/v1")

	// Mock auth middleware
	mockAuthMiddleware := func(c *gin.Context) {
		c.Next()
	}

	memberGroup := NewMemberRouter(baseGroup, mockAuthMiddleware)

	// Assertions
	assert.NotNil(t, memberGroup, "Router group should not be nil")

	// Verify routes are set up correctly
	routes := router.Routes()

	// Check if POST /members/ endpoint exists
	foundMembersRoute := false
	foundMembersListRoute := false
	for _, route := range routes {
		if route.Method == "POST" && route.Path == "/api/v1/members/" {
			foundMembersRoute = true
		}
		if route.Method == "POST" && route.Path == "/api/v1/members/list" {
			foundMembersListRoute = true
		}
	}
	assert.True(t, foundMembersRoute, "POST /members/ route should exist")
	assert.True(t, foundMembersListRoute, "POST /members/list route should exist")
}
