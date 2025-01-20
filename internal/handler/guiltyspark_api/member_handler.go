package handler

import (
	"fmt"
	"net/http"

	service "github.com/cbglabs/nexusgate-pbmapi/internal/service/guiltyspark_api"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"
	"github.com/gin-gonic/gin"
)

// MemberHandler defines the interface for member operations
type MemberHandler interface {
	ListMembers(c *gin.Context)
	CheckEligibility(c *gin.Context)
}

// memberHandler implements the MemberHandler interface
type memberHandler struct {
	service service.MemberService
}

// Verify memberHandler implements MemberHandler
var _ MemberHandler = (*memberHandler)(nil)

// NewMemberHandler creates a new MemberHandler instance
func NewMemberHandler(service service.MemberService) MemberHandler {
	return &memberHandler{
		service: service,
	}
}

// ListMembers handles the query of members
func (h *memberHandler) ListMembers(c *gin.Context) {
	var variables types.MemberQueryVariables
	if err := c.ShouldBindJSON(&variables); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Data:   fmt.Sprintf("Invalid request, missing required fields: %v", err),
			Status: "fail",
		})
		return
	}

	result, err := h.service.ListMembers(c, variables)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Data:   fmt.Sprintf("Failed to query members: %v", err),
			Status: "fail",
		})
		return
	}

	// Respond with a success message.
	c.JSON(http.StatusOK, response.APIResponse{
		Data:   result,
		Status: "success",
	})
}

// CheckEligibility handles the query of members to check if they are eligible
func (h *memberHandler) CheckEligibility(c *gin.Context) {
	var variables types.MemberQueryVariables
	if err := c.ShouldBindJSON(&variables); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Data:   fmt.Sprintf("Invalid request, missing required fields: %v", err),
			Status: "fail",
		})
		return
	}

	result, err := h.service.CheckEligibility(c, variables)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Data:   fmt.Sprintf("Failed to query members: %v", err),
			Status: "fail",
		})
		return
	}

	// Respond with a success message.
	c.JSON(http.StatusOK, response.APIResponse{
		Data:   result,
		Status: "success",
	})
}
