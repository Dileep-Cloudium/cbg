package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cbglabs/nexusgate-pbmapi/internal/db"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"
)

// MemberService defines the interface for member operations
type MemberService interface {
	ListMembers(ctx context.Context, variables types.MemberQueryVariables) (*response.MemberResponse, error)
	CheckEligibility(ctx context.Context, variables types.MemberQueryVariables) (bool, error)
}

// memberService implements the MemberService interface
type memberService struct {
	dbClient db.GraphQLClient
}

// Verify memberService implements MemberService
var _ MemberService = (*memberService)(nil)

// NewMemberService creates a new MemberService instance
func NewMemberService(dbClient db.GraphQLClient) MemberService {
	return &memberService{
		dbClient: dbClient,
	}
}

// ListMembers handles the query of members
func (s *memberService) ListMembers(ctx context.Context, variables types.MemberQueryVariables) (*response.MemberResponse, error) {
	const query = `
			query QueryMembers(
				$customerCode: String,
				$clientCode: String,
				$groupCode: String,
				$memberId: String,
				$personCode: String
			) {
				master_elig_cache(
					where: {
						customer_code: { _eq: $customerCode }
						client_code: { _eq: $clientCode }
						group_code: { _eq: $groupCode }
						member_id: { _eq: $memberId }
						person_code: { _eq: $personCode }
					}
				) {
					jv_partner_name
					tpa_name
					customer_code
					customer_name
					client_code
					client_name
					group_code
					group_name
					internal_unique_id
					insured_id
					member_id
					person_code
					extended_member_id
					alternate_member_id
					first_name
					middle_name
					last_name
					gender
					address1
					address2
					address3
					city
					state
					zip_code
					birth_date
					effective_date
					termination_date
					primary_email
					primary_phone
				}
			}
	`

	// Convert struct to map using JSON marshal/unmarshal
	jsonData, err := json.Marshal(variables)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal variables: %w", err)
	}

	// Convert JSON to map[string]interface{} for GraphQL
	variablesMap := make(map[string]interface{})
	if err := json.Unmarshal(jsonData, &variablesMap); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal variables: %w", err)
	}

	req := &types.GraphQLRequest{
		Query:     query,
		Variables: variablesMap,
	}

	var resp types.MemberGraphQLResponse
	if err := s.dbClient.Execute(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &response.MemberResponse{
		Members: resp.Members,
	}, nil
}

// CheckEligibility checks if a member has active coverage based on effective and termination dates
func (s *memberService) CheckEligibility(ctx context.Context, variables types.MemberQueryVariables) (bool, error) {
	resp, err := s.ListMembers(ctx, variables)
	if err != nil {
		return false, fmt.Errorf("Failed to get member: %w", err)
	}

	if len(resp.Members) == 0 {
		return false, nil
	}

	member := resp.Members[0]
	now := time.Now()

	// Use the Time values directly from CustomTime
	effectiveDate := member.EffectiveDate.Time
	terminationDate := member.TerminationDate.Time

	// Check if current date falls within eligibility period
	return !now.Before(effectiveDate) && !now.After(terminationDate), nil
}
