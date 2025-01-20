package handler

import (
	"fmt"
	"net/http"

	service "github.com/cbglabs/nexusgate-pbmapi/internal/service/portal_api"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/request"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"

	"github.com/gin-gonic/gin"
)

// GetStates retrieves the list of all states.
// It calls the service layer to get the list and return it to the client.
// @Summary Get list of states
// @Description Retrieve the list of all states from the system.
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/public/states [get]
func GetStates(c *gin.Context) {
	statesList, err := service.GetStates()

	// Log and respond with an error if the service call fails.
	if err != nil {
		fmt.Println("error occurred")
		c.JSON(http.StatusInternalServerError, gin.H{"Data": "Failed to retrieve states", "Status": "fail"})
		return
	}

	// Respond with the list of states and a success message.
	c.JSON(200, response.APIResponse{
		Data:   statesList,
		Status: "success",
	})
}

// Login validates the user and generates token.
// It calls the service layer to generate token.
// @Summary Validate user and token generation
// @Description Validate the username and generate a token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param 	request body request.LoginRequest true "Login request"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/auth/login [post]
func Login(c *gin.Context) {
	var loginRequest request.LoginRequest

	// Bind JSON input to loginRequest. Respond with a failure if the JSON is invalid.
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(200, response.APIResponse{
			Data:   err.Error(),
			Status: "fail",
		})
		return
	}

	loginData, err := service.Login(loginRequest)

	// Log and respond with an error if the service call fails.
	if err != nil {
		fmt.Println("error occurred")
		c.JSON(http.StatusInternalServerError, gin.H{"Data": "Failed to login", "Status": "fail"})
		return
	}

	var loginResponse response.LoginResponse
	loginResponse.Token = loginData

	// Respond with the token and a success message.
	c.JSON(200, response.APIResponse{
		Data:   loginResponse,
		Status: "success",
	})
}

// Registration of a user.
// It calls the service layer to register a user.
// @Summary User Registration
// @Description A user will be registered for authentication.
// @Tags Auth
// @Accept json
// @Produce json
// @Param 	request body request.RegisterRequest true "Register request"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/auth/register [post]
func Register(c *gin.Context) {
	var registerRequest request.RegisterRequest

	// Bind JSON input to registerRequest. Respond with a failure if the JSON is invalid.
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(200, response.APIResponse{
			Data:   err.Error(),
			Status: "fail",
		})
		return
	}

	// Call the service layer to add the user.
	id := service.Register(registerRequest)

	var registerResponse response.RegisterResponse
	registerResponse.Id = id
	registerResponse.Message = "User Successfully Registered"

	// Respond with a success message.
	c.JSON(200, response.APIResponse{
		Data:   registerResponse,
		Status: "success",
	})
}
