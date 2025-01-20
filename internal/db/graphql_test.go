package db

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cbglabs/nexusgate-pbmapi/internal/types"
)

func TestNewGraphQLClient(t *testing.T) {
	tests := []struct {
		name        string
		endpoint    string
		apiKey      string
		wantErr     bool
		errContains string
	}{
		{
			name:        "missing endpoint",
			endpoint:    "",
			apiKey:      "test-key",
			wantErr:     true,
			errContains: "GraphQL endpoint is not set",
		},
		{
			name:        "missing api key",
			endpoint:    "http://test.com",
			apiKey:      "",
			wantErr:     true,
			errContains: "GraphQL API key is not set",
		},
		{
			name:     "successful creation",
			endpoint: "http://test.com",
			apiKey:   "test-key",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("GRAPHQL_ENDPOINT", tt.endpoint)
			os.Setenv("GRAPHQL_API_KEY", tt.apiKey)
			defer os.Unsetenv("GRAPHQL_ENDPOINT")
			defer os.Unsetenv("GRAPHQL_API_KEY")

			client, err := NewGraphQLClient()
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				if tt.errContains != "" && err.Error() != tt.errContains {
					t.Errorf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if client == nil {
				t.Error("expected client, got nil")
			}
		})
	}
}

func TestGraphQLClient_Execute(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("expected Content-Type header to be application/json")
		}
		if r.Header.Get("x-hasura-admin-secret") != "test-key" {
			t.Error("expected x-hasura-admin-secret header to be test-key")
		}

		// Return mock response
		resp := types.GraphQLResponse{
			Data: map[string]interface{}{
				"test": "data",
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Set up client
	os.Setenv("GRAPHQL_ENDPOINT", server.URL)
	os.Setenv("GRAPHQL_API_KEY", "test-key")
	defer os.Unsetenv("GRAPHQL_ENDPOINT")
	defer os.Unsetenv("GRAPHQL_API_KEY")

	client, err := NewGraphQLClient()
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// Test successful request
	req := &types.GraphQLRequest{
		Query: "query { test }",
	}
	var resp map[string]interface{}

	err = client.Execute(context.Background(), req, &resp)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp["test"] != "data" {
		t.Errorf("expected response data to be 'data', got %v", resp["test"])
	}
}
