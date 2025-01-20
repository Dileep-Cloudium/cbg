package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewHealthRouter(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()

	// Create the base router group
	apiGroup := router.Group("/api/v1")

	// Set up the health routes
	NewHealthRouter(apiGroup)

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert no error in parsing
	assert.NoError(t, err)

	// Assert response content
	assert.Equal(t, "ok", response["status"])
}
