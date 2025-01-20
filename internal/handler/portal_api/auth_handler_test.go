package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"testing"

	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserRegister(t *testing.T) {
	os.Setenv("AWS_REGION", "us-west-2")
	os.Setenv("DYNAMODB_TABLE_NAME", "NexusGate-PBMAPI-AllItems1")
	os.Setenv("AWS_PROFILE", "arcadia-develop")

	gin.SetMode(gin.TestMode)
	baseRouter := gin.Default()

	baseRouter.POST("/api/v1/auth/register", Register)

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedMsg    string
	}{
		{
			name: "Valid request with first_name, last_name, email and password",
			requestBody: `{"first_name": "Test first name","last_name":"Test last name",
				"email":"testsowjanya@gmail.com","password":"test password"}`,
			expectedStatus: http.StatusOK,
			expectedMsg:    "User Successfully Registered",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create fresh recorder and context for each test
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			Register(c)

			var response response.APIResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

// func TestLogin(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	tests := []struct {
// 		name           string
// 		requestBody    string
// 		expectedStatus int
// 	}{
// 		{
// 			name:           "Valid request with first_name, last_name, email and password",
// 			requestBody:    `{"email":"testsowjanya@gmail.com","password":"test password"}`,
// 			expectedStatus: http.StatusOK,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create fresh recorder and context for each test
// 			w := httptest.NewRecorder()

// 			c, _ := gin.CreateTestContext(w)

// 			req := httptest.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(tt.requestBody))
// 			req.Header.Set("Content-Type", "application/json")
// 			c.Request = req

// 			Login(c)

// 			var response response.APIResponse
// 			err := json.Unmarshal(w.Body.Bytes(), &response)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expectedStatus, w.Code)
// 		})
// 	}
// }

func TestGetStates(t *testing.T) {
	gin.SetMode(gin.TestMode)
	baseRouter := gin.Default()

	baseRouter.GET("/api/v1/public/states", GetStates)
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Get states",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create fresh recorder and context for each test
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("GET", "/api/v1/public/states", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			GetStates(c)

			var response response.APIResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
