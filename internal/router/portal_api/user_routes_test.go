package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mockAuthMiddleware creates a test authentication middleware
func mockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simulate authentication by setting a user ID in the context
		c.Set("userID", "test-user-id")
		c.Next()
	}
}

// mockFailedAuthMiddleware creates a test middleware that simulates authentication failure
func mockFailedAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}

func TestNewUserRouter(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	validUserPayload := map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"email":      "john.doe@example.com",
		"password":   "Password123",
	}

	validUpdatePayload := map[string]interface{}{
		"first_name": "Johnny",
		"last_name":  "Doe",
		"email":      "johnny.doe@example.com",
	}

	jsonPayload, _ := json.Marshal(validUserPayload)
	updatePayload, _ := json.Marshal(validUpdatePayload)

	tests := []struct {
		name           string
		method         string
		url            string
		body           []byte
		authMiddleware gin.HandlerFunc
		expectedCode   int
	}{
		{
			name:           "Add User - Authorized",
			method:         "POST",
			url:            "/api/users/",
			body:           jsonPayload,
			authMiddleware: mockAuthMiddleware(),
			expectedCode:   http.StatusOK,
		},
		{
			name:           "Add User - Unauthorized",
			method:         "POST",
			url:            "/api/users/",
			body:           jsonPayload,
			authMiddleware: mockFailedAuthMiddleware(),
			expectedCode:   http.StatusUnauthorized,
		},
		{
			name:           "Get Users - Authorized",
			method:         "GET",
			url:            "/api/users/list",
			authMiddleware: mockAuthMiddleware(),
			expectedCode:   http.StatusOK,
		},
		{
			name:           "Update User - Unauthorized",
			method:         "PUT",
			url:            "/api/users/somer-user-id",
			body:           updatePayload,
			authMiddleware: mockFailedAuthMiddleware(),
			expectedCode:   http.StatusUnauthorized,
		},
		// {
		// 	name:           "Delete User - Authorized",
		// 	method:         "DELETE",
		// 	url:            "/api/users/test-user-id",
		// 	authMiddleware: mockAuthMiddleware(),
		// 	expectedCode:   http.StatusOK,
		// },
		{
			name:           "Delete User - Unauthorized",
			method:         "DELETE",
			url:            "/api/users/test-user-id",
			authMiddleware: mockFailedAuthMiddleware(),
			expectedCode:   http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Gin router
			router := gin.New()

			// Create the base router group
			apiGroup := router.Group("/api")

			// Set up the user routes with the test middleware
			NewUserRouter(apiGroup, tt.authMiddleware)

			// Create a test request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(tt.body))
			req.Header.Set("Content-Type", "application/json")

			// Serve the request
			router.ServeHTTP(w, req)

			// Assert the response
			if !assert.Equal(t, tt.expectedCode, w.Code) {
				t.Errorf("Response body: %s", w.Body.String())
			}
		})
	}
}
