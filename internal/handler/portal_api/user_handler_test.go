package handler

import (
	// "encoding/json"

	"net/http"
	"net/http/httptest"
	"os"

	// "strings"
	"testing"

	// auth "github.com/cbglabs/nexusgate-pbmapi/internal/middleware"

	"github.com/cbglabs/nexusgate-pbmapi/internal/db"
	service "github.com/cbglabs/nexusgate-pbmapi/internal/service/portal_api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func MockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simulate authentication by setting a test user in context
		c.Set("user", "test_user")
		c.Next()
	}
}

func TestListUsers(t *testing.T) {
	os.Setenv("AWS_REGION", "us-west-2")
	os.Setenv("DYNAMODB_TABLE_NAME", "NexusGate-PBMAPI-AllItems1")
	os.Setenv("AWS_PROFILE", "arcadia-develop")

	dbClient, _ := db.NewDynamoDBClient()
	userService := service.NewUserService(dbClient)
	userHandler := NewUserHandler(userService)

	gin.SetMode(gin.TestMode)
	baseRouter := gin.Default()

	userGroup := baseRouter.Group("/api/v1/user")
	userGroup.Use(MockAuthMiddleware())
	userGroup.GET("/list", userHandler.ListUsers)

	// Create a test HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/list", nil)

	// Use the Gin engine (baseRouter) to handle the HTTP request
	baseRouter.ServeHTTP(w, req)
	// Assert the expected HTTP status code and response body
	assert.Equal(t, 200, w.Code)
}

// func TestCreateUser(t *testing.T) {
// 	// Setup
// 	gin.SetMode(gin.TestMode)

// 	tests := []struct {
// 		name           string
// 		requestBody    string
// 		expectedStatus int
// 		expectedMsg    string
// 	}{
// 		{
// 			name: "Valid request with first_name, last_name, email and password",
// 			requestBody: `{"first_name": "Test first name","last_name":"Test last name",
// 				"email":"testsowjanya@gmail.com","password":"test password"}`,
// 			expectedStatus: http.StatusOK,
// 			expectedMsg:    "Successfully Added user.",
// 		},
// 		{
// 			name: "Invalid JSON",
// 			requestBody: `{"first_name": "Test first name","last_name":"Test last name",
// 				"email":"testgmailcom","password":"testpassword"}`,
// 			expectedStatus: http.StatusOK,
// 			expectedMsg:    "Key: 'AddUserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create fresh recorder and context for each test
// 			w := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(w)

// 			req := httptest.NewRequest("POST", "/api/v1/user", strings.NewReader(tt.requestBody))
// 			req.Header.Set("Content-Type", "application/json")
// 			c.Request = req

// 			AddUser(c)

// 			var response response.APIResponse
// 			err := json.Unmarshal(w.Body.Bytes(), &response)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expectedStatus, w.Code)
// 		})
// 	}
// }

// func TestGetUser(t *testing.T) {
// 	// Setup
// 	gin.SetMode(gin.TestMode)

// 	tests := []struct {
// 		name           string
// 		requestParam   string
// 		expectedStatus int
// 		expectedMsg    interface{}
// 	}{
// 		{
// 			name:           "Get User Success",
// 			requestParam:   "2088b8da-98d8-45be-bcad-47a39d6c67d7",
// 			expectedStatus: http.StatusOK,
// 			expectedMsg:    response.APIResponse{},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create fresh recorder and context for each test
// 			w := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(w)

// 			req := httptest.NewRequest("GET", "/api/v1/user"+tt.requestParam, nil)
// 			req.Header.Set("Content-Type", "application/json")
// 			c.Request = req

// 			GetUser(c)

// 			var response response.APIResponse
// 			err := json.Unmarshal(w.Body.Bytes(), &response)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expectedStatus, w.Code)
// 		})
// 	}
// }

// func TestUpdateUser(t *testing.T) {
// 	// Setup
// 	gin.SetMode(gin.TestMode)

// 	tests := []struct {
// 		name           string
// 		requestParam   string
// 		requestBody    string
// 		expectedStatus int
// 		expectedMsg    string
// 	}{
// 		{
// 			name:           "Update User Success",
// 			requestParam:   "2088b8da-98d8-45be-bcad-47a39d6c67d7",
// 			requestBody:   `{"first_name": "Test first name","last_name":"Test last name"}`,
// 			expectedStatus: http.StatusOK,
// 			expectedMsg:    "User Updated Successfully.",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create fresh recorder and context for each test
// 			w := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(w)

// 			req := httptest.NewRequest("PUT", "/api/v1/user"+tt.requestParam, strings.NewReader(tt.requestBody))
// 			req.Header.Set("Content-Type", "application/json")
// 			c.Request = req

// 			UpdateUser(c)

// 			var response response.APIResponse
// 			err := json.Unmarshal(w.Body.Bytes(), &response)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expectedStatus, w.Code)
// 		})
// 	}
// }

// func TestDeleteUser(t *testing.T) {
// 	// Setup
// 	gin.SetMode(gin.TestMode)

// 	tests := []struct {
// 		name           string
// 		requestParam   string
// 		requestBody    string
// 		expectedStatus int
// 		expectedMsg    string
// 	}{
// 		{
// 			name:           "Delete User Success",
// 			requestParam:   "12aff201-e1f2-490d-ab86-ae0c144c9e47",
// 			expectedStatus: http.StatusOK,
// 			expectedMsg:    "User Deleted Successfully.",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create fresh recorder and context for each test
// 			w := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(w)

// 			req := httptest.NewRequest("DELETE", "/api/v1/user"+tt.requestParam, nil)
// 			req.Header.Set("Content-Type", "application/json")
// 			c.Request = req

// 			DeleteUser(c)

// 			var response response.APIResponse
// 			err := json.Unmarshal(w.Body.Bytes(), &response)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expectedStatus, w.Code)
// 		})
// 	}
// }
