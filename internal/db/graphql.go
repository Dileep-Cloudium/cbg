package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cbglabs/nexusgate-pbmapi/internal/types"
)

// GraphQLClient defines the interface for GraphQL operations
type GraphQLClient interface {
	Execute(ctx context.Context, req *types.GraphQLRequest, resp interface{}) error
}

// graphQLClient implements the GraphQLClient interface
type graphQLClient struct {
	endpoint string
	apiKey   string
}

// Verify graphQLClient implements GraphQLClient
var _ GraphQLClient = (*graphQLClient)(nil)

// NewGraphQLClient creates a new GraphQL client
func NewGraphQLClient() (GraphQLClient, error) {
	endpoint := os.Getenv("GRAPHQL_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("GraphQL endpoint is not set")
	}

	apiKey := os.Getenv("GRAPHQL_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GraphQL API key is not set")
	}

	return &graphQLClient{
		endpoint: endpoint,
		apiKey:   apiKey,
	}, nil
}

// Execute sends a GraphQL request to the server and returns the response
func (c *graphQLClient) Execute(ctx context.Context, req *types.GraphQLRequest, resp interface{}) error {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-hasura-admin-secret", c.apiKey)

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	var graphQLResp types.GraphQLResponse
	graphQLResp.Data = resp

	if err := json.NewDecoder(httpResp.Body).Decode(&graphQLResp); err != nil {
		return err
	}

	if len(graphQLResp.Errors) > 0 {
		return fmt.Errorf("GraphQL errors: %v", graphQLResp.Errors)
	}

	return nil
}
