package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogger(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Default().Writer())

	// Set up Gin router with logger middleware
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(Logger())

	// Add a test endpoint
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Create test request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check response status
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Verify log output contains expected information
	logOutput := buf.String()
	expectedParts := []string{
		"Request: GET /test",
		"Status: 200",
		"Duration:",
	}

	for _, part := range expectedParts {
		if !strings.Contains(logOutput, part) {
			t.Errorf("Expected log to contain '%s', got: %s", part, logOutput)
		}
	}
}
