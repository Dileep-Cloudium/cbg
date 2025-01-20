package service

import (
	"context"
	"testing"
	"time"

	"github.com/cbglabs/nexusgate-pbmapi/internal/types"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockGraphQLClient is a mock implementation of the GraphQLClient interface
type MockGraphQLClient struct {
	mock.Mock
}

func (m *MockGraphQLClient) Execute(ctx context.Context, req *types.GraphQLRequest, resp interface{}) error {
	args := m.Called(ctx, req, resp)
	return args.Error(0)
}

func TestListMembers(t *testing.T) {
	tests := []struct {
		name      string
		variables types.MemberQueryVariables
		mockResp  types.MemberGraphQLResponse
		wantErr   bool
	}{
		{
			name: "successful query",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				MemberID:     "MEM456",
			},
			mockResp: types.MemberGraphQLResponse{
				Members: []response.Member{
					{
						CustomerCode: "CUST123",
						MemberID:     "MEM456",
						FirstName:    "John",
						LastName:     "Doe",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := new(MockGraphQLClient)

			// Set up mock expectation
			mockClient.On("Execute",
				mock.Anything,
				mock.MatchedBy(func(req *types.GraphQLRequest) bool {
					// Basic validation that the request contains required fields
					return req.Query != "" && req.Variables != nil
				}),
				mock.AnythingOfType("*types.MemberGraphQLResponse"),
			).Run(func(args mock.Arguments) {
				// Set the mock response
				resp := args.Get(2).(*types.MemberGraphQLResponse)
				*resp = tt.mockResp
			}).Return(nil)

			// Create service with mock client
			svc := NewMemberService(mockClient)

			// Execute test
			got, err := svc.ListMembers(context.Background(), tt.variables)

			// Verify results
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.mockResp.Members, got.Members)
			}

			// Verify mock expectations
			mockClient.AssertExpectations(t)
		})
	}
}

func TestCheckEligibility(t *testing.T) {
	tests := []struct {
		name      string
		variables types.MemberQueryVariables
		mockResp  types.MemberGraphQLResponse
		wantErr   bool
		want      bool
	}{
		{
			name: "eligible member - within date range",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				ClientCode:   "CLIENT123",
				GroupCode:    "GROUP123",
				MemberID:     "MEMBER123",
				PersonCode:   "PERSON123",
			},
			mockResp: types.MemberGraphQLResponse{
				Members: []response.Member{
					{
						CustomerCode:    "CUST123",
						ClientCode:      "CLIENT123",
						GroupCode:       "GROUP123",
						MemberID:        "MEMBER123",
						PersonCode:      "PERSON123",
						EffectiveDate:   response.CustomTime{Time: time.Now().AddDate(0, -6, 0).UTC()},
						TerminationDate: response.CustomTime{Time: time.Now().AddDate(0, 6, 0).UTC()},
					},
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ineligible member - future effective date",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				ClientCode:   "CLIENT123",
				GroupCode:    "GROUP123",
				MemberID:     "MEMBER123",
				PersonCode:   "PERSON123",
			},
			mockResp: types.MemberGraphQLResponse{
				Members: []response.Member{
					{
						CustomerCode:    "CUST123",
						ClientCode:      "CLIENT123",
						GroupCode:       "GROUP123",
						MemberID:        "MEMBER123",
						PersonCode:      "PERSON123",
						EffectiveDate:   response.CustomTime{Time: time.Now().AddDate(0, 6, 0).UTC()},
						TerminationDate: response.CustomTime{Time: time.Now().AddDate(0, 12, 0).UTC()},
					},
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "ineligible member - past termination date",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				ClientCode:   "CLIENT123",
				GroupCode:    "GROUP123",
				MemberID:     "MEMBER123",
				PersonCode:   "PERSON123",
			},
			mockResp: types.MemberGraphQLResponse{
				Members: []response.Member{
					{
						CustomerCode:    "CUST123",
						ClientCode:      "CLIENT123",
						GroupCode:       "GROUP123",
						MemberID:        "MEMBER123",
						PersonCode:      "PERSON123",
						EffectiveDate:   response.CustomTime{Time: time.Now().AddDate(0, -6, 0).UTC()},
						TerminationDate: response.CustomTime{Time: time.Now().AddDate(0, -1, 0).UTC()},
					},
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "member not found",
			variables: types.MemberQueryVariables{
				CustomerCode: "CUST123",
				ClientCode:   "CLIENT123",
				GroupCode:    "GROUP123",
				MemberID:     "MEMBER123",
				PersonCode:   "PERSON123",
			},
			mockResp: types.MemberGraphQLResponse{
				Members: []response.Member{},
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := new(MockGraphQLClient)

			// Set up mock expectation
			mockClient.On("Execute",
				mock.Anything,
				mock.MatchedBy(func(req *types.GraphQLRequest) bool {
					return req.Query != "" && req.Variables != nil
				}),
				mock.AnythingOfType("*types.MemberGraphQLResponse"),
			).Run(func(args mock.Arguments) {
				resp := args.Get(2).(*types.MemberGraphQLResponse)
				*resp = tt.mockResp
			}).Return(nil)

			// Create service with mock client
			svc := NewMemberService(mockClient)

			// Execute test
			got, err := svc.CheckEligibility(context.Background(), tt.variables)

			// Verify results
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			// Verify mock expectations
			mockClient.AssertExpectations(t)
		})
	}
}
