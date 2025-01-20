package response

import (
	"fmt"
	"strings"
	"time"
)

type MemberResponse struct {
	Members []Member `json:"members"`
}

// CustomTime is a wrapper around time.Time that implements custom JSON unmarshaling
type CustomTime struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler interface
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	// Try parsing with different formats
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		// Fallback to RFC3339 format if simple date format fails
		t, err = time.Parse(time.RFC3339, s)
		if err != nil {
			return err
		}
	}
	ct.Time = t
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format("2006-01-02"))), nil
}

// MemberGraphQLData represents a single member's data from the GraphQL service
type Member struct {
	JVPartnerName     string     `json:"jv_partner_name"`
	TPAName           string     `json:"tpa_name"`
	CustomerCode      string     `json:"customer_code"`
	CustomerName      string     `json:"customer_name"`
	ClientCode        string     `json:"client_code"`
	ClientName        string     `json:"client_name"`
	GroupCode         string     `json:"group_code"`
	GroupName         string     `json:"group_name"`
	InternalUniqueID  int        `json:"internal_unique_id"`
	InsuredID         string     `json:"insured_id"`
	MemberID          string     `json:"member_id"`
	PersonCode        string     `json:"person_code"`
	ExtendedMemberID  string     `json:"extended_member_id"`
	AlternateMemberID string     `json:"alternate_member_id"`
	FirstName         string     `json:"first_name"`
	MiddleName        string     `json:"middle_name"`
	LastName          string     `json:"last_name"`
	Gender            string     `json:"gender"`
	Address1          string     `json:"address1"`
	Address2          string     `json:"address2"`
	Address3          string     `json:"address3"`
	City              string     `json:"city"`
	State             string     `json:"state"`
	ZipCode           string     `json:"zip_code"`
	BirthDate         CustomTime `json:"birth_date"`
	EffectiveDate     CustomTime `json:"effective_date"`
	TerminationDate   CustomTime `json:"termination_date"`
	PrimaryEmail      string     `json:"primary_email"`
	PrimaryPhone      string     `json:"primary_phone"`
}
