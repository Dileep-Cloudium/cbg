package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TestAuthSimpleMiddleware(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		authHeader string
		wantCode   int
	}{
		{
			name:       "Valid API Key",
			authHeader: "Bearer ng-api-key-f7d9a2e1b3c5",
			wantCode:   http.StatusOK,
		},
		{
			name:       "Invalid API Key",
			authHeader: "Bearer invalid-key",
			wantCode:   http.StatusUnauthorized,
		},
		{
			name:       "Missing Bearer Prefix",
			authHeader: "ng-api-key-f7d9a2e1b3c5",
			wantCode:   http.StatusUnauthorized,
		},
		{
			name:       "Empty Authorization",
			authHeader: "",
			wantCode:   http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create test request
			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", tt.authHeader)
			c.Request = req

			// Test middleware
			AuthSimpleMiddleware()(c)

			if w.Code != tt.wantCode {
				t.Errorf("AuthSimpleMiddleware() status code = %v, want %v", w.Code, tt.wantCode)
			}
		})
	}
}

func TestAuthMiddleware(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	os.Setenv("SECRET_KEY", "test-secret-key")

	// Create a valid JWT token for testing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": "test-user-id",
	})
	validTokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	tests := []struct {
		name       string
		authHeader string
		wantCode   int
	}{
		{
			name:       "Valid JWT",
			authHeader: validTokenString,
			wantCode:   http.StatusOK,
		},
		{
			name:       "Invalid JWT",
			authHeader: "invalid.token.string",
			wantCode:   http.StatusUnauthorized,
		},
		{
			name:       "Empty Authorization",
			authHeader: "",
			wantCode:   http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create test request
			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", tt.authHeader)
			c.Request = req

			// Test middleware
			AuthMiddleware()(c)

			if w.Code != tt.wantCode {
				t.Errorf("AuthMiddleware() status code = %v, want %v", w.Code, tt.wantCode)
			}
		})
	}
}
