package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	// Create a new router
	router := gin.New()
	return router
}

func TestNewAuthRouter(t *testing.T) {
	// Setup
	router := setupTestRouter()
	routerGroup := router.Group("/api")

	// Execute
	NewAuthRouter(routerGroup)

	// Test cases
	tests := []struct {
		name     string
		method   string
		path     string
		expected int
	}{
		{
			name:     "Login route should exist",
			method:   http.MethodPost,
			path:     "/api/auth/login",
			expected: http.StatusOK,
		},
		{
			name:     "Register route should exist",
			method:   http.MethodPost,
			path:     "/api/auth/register",
			expected: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			router.ServeHTTP(w, req)

			// Assert that the route exists (will return 200 if handler exists, 404 if not)
			assert.NotEqual(t, http.StatusNotFound, w.Code, "Route should exist")
		})
	}
}

func TestNewPublicRouter(t *testing.T) {
	// Setup
	router := setupTestRouter()
	routerGroup := router.Group("/api")

	// Execute
	NewPublicRouter(routerGroup)

	// Test cases
	tests := []struct {
		name     string
		method   string
		path     string
		expected int
	}{
		{
			name:     "Get states route should exist",
			method:   http.MethodGet,
			path:     "/api/public/states",
			expected: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			router.ServeHTTP(w, req)

			// Assert that the route exists (will return 200 if handler exists, 404 if not)
			assert.NotEqual(t, http.StatusNotFound, w.Code, "Route should exist")
		})
	}
}
