package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cbglabs/nexusgate-pbmapi/internal/types"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockMemberService is a mock implementation of MemberService
type MockMemberService struct {
	mock.Mock
}

func (m *MockMemberService) ListMembers(c context.Context, variables types.MemberQueryVariables) (*response.MemberResponse, error) {
	args := m.Called(c, variables)
	return args.Get(0).(*response.MemberResponse), args.Error(1)
}

func (m *MockMemberService) CheckEligibility(c context.Context, variables types.MemberQueryVariables) (bool, error) {
	args := m.Called(c, variables)
	return args.Bool(0), args.Error(1)
}

func TestListMembers(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		variables      types.MemberQueryVariables
		mockReturn     *response.MemberResponse
		mockError      error
		expectedStatus int
	}{
		{
			name: "Success case",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				ClientCode:   "CLIENT123",
				GroupCode:    "GROUP123",
				MemberID:     "MEMBER123",
				PersonCode:   "PERSON123",
			},
			mockReturn:     &response.MemberResponse{Members: []response.Member{{MemberID: "member1"}, {MemberID: "member2"}}},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name: "Service error case",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				ClientCode:   "CLIENT123",
				GroupCode:    "GROUP123",
				MemberID:     "MEMBER123",
				PersonCode:   "PERSON123",
			},
			mockReturn:     nil,
			mockError:      fmt.Errorf("service error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockMemberService)
			mockService.On("ListMembers", mock.Anything, tt.variables).Return(tt.mockReturn, tt.mockError)

			// Create handler with mock service
			handler := NewMemberHandler(mockService)

			// Create test context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create request with JSON body
			jsonBody, _ := json.Marshal(tt.variables)
			c.Request = httptest.NewRequest(http.MethodPost, "/members", bytes.NewBuffer(jsonBody))
			c.Request.Header.Set("Content-Type", "application/json")

			// Call handler
			handler.ListMembers(c)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestCheckEligibility(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		variables      types.MemberQueryVariables
		mockReturn     bool
		mockError      error
		expectedStatus int
	}{
		{
			name: "Success case",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				ClientCode:   "CLIENT123",
				GroupCode:    "GROUP123",
				MemberID:     "MEMBER123",
				PersonCode:   "PERSON123",
			},
			mockReturn:     true,
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name: "Service error case",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				ClientCode:   "CLIENT123",
				GroupCode:    "GROUP123",
				MemberID:     "MEMBER123",
				PersonCode:   "PERSON123",
			},
			mockReturn:     false,
			mockError:      fmt.Errorf("service error"),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Invalid request case",
			variables:      types.MemberQueryVariables{}, // Empty variables to trigger validation error
			mockReturn:     false,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockMemberService)

			// Only set up mock expectations if we expect the service to be called
			if tt.variables != (types.MemberQueryVariables{}) {
				mockService.On("CheckEligibility", mock.Anything, tt.variables).Return(tt.mockReturn, tt.mockError)
			}

			// Create handler with mock service
			handler := NewMemberHandler(mockService)

			// Create test context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create request with JSON body
			jsonBody, _ := json.Marshal(tt.variables)
			c.Request = httptest.NewRequest(http.MethodPost, "/members/check-eligibility", bytes.NewBuffer(jsonBody))
			c.Request.Header.Set("Content-Type", "application/json")

			// Call handler
			handler.CheckEligibility(c)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)

			if tt.expectedStatus == http.StatusOK {
				var response response.APIResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, "success", response.Status)
				assert.NotNil(t, response.Data)
			}
		})
	}
}
