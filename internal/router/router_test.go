package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMountRoutes(t *testing.T) {
	// Initialize router
	r := MountRoutes()

	t.Run("Home route", func(t *testing.T) {
		// Create a new HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		// Serve the request
		r.ServeHTTP(w, req)

		// Assert response
		assert.Equal(t, 200, w.Code)

		// Check response body
		var response string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Welcome home", response)
	})

	t.Run("Health check route", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/portal-api/health", nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Router groups are properly mounted", func(t *testing.T) {
		// Test that the router has the expected group paths
		routes := r.Routes()

		// Helper function to check if a path exists
		hasPath := func(path string) bool {
			for _, route := range routes {
				if route.Path == path {
					return true
				}
			}
			return false
		}

		// Check for essential paths
		assert.True(t, hasPath("/"))
		assert.True(t, hasPath("/swagger/*any"))
		assert.True(t, hasPath("/api/v1/portal-api/health"))
	})
}
