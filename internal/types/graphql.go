package types

import "github.com/cbglabs/nexusgate-pbmapi/internal/types/response"

// GraphQLRequest represents a generic GraphQL request structure
type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
}

// GraphQLResponse represents a generic GraphQL response structure
type GraphQLResponse struct {
	Data   interface{}    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

// GraphQLError represents a GraphQL error
type GraphQLError struct {
	Message    string                 `json:"message"`
	Locations  []GraphQLErrorLocation `json:"locations,omitempty"`
	Path       []interface{}          `json:"path,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

type GraphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// MemberQueryVariables represents the variables for member queries
type MemberQueryVariables struct {
	CustomerCode string `json:"customerCode" binding:"required"`
	ClientCode   string `json:"clientCode" binding:"required"`
	GroupCode    string `json:"groupCode" binding:"required"`
	MemberID     string `json:"memberId" binding:"required"`
	PersonCode   string `json:"personCode" binding:"required"`
}

// MemberGraphQLResponse represents the raw response from the GraphQL service
type MemberGraphQLResponse struct {
	// The field name matches the GraphQL response exactly
	Members []response.Member `json:"master_elig_cache"`
}
