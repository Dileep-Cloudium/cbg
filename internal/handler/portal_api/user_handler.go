package handler

import (
	"fmt"
	"net/http"

	service "github.com/cbglabs/nexusgate-pbmapi/internal/service/portal_api"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/request"
	"github.com/cbglabs/nexusgate-pbmapi/internal/types/response"

	"github.com/gin-gonic/gin"
)

// UserHandler defines the interface for user operations
type UserHandler interface {
	CreateUser(c *gin.Context)
	ReadUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	ListUsers(c *gin.Context)
}

// userHandler implements the UserHandler interface
type userHandler struct {
	service service.UserService
}

// Verify userHandler implements UserHandler
var _ UserHandler = (*userHandler)(nil)

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		service: userService,
	}
}

// AddUser handles the addition of a new user.
// It binds JSON input to a request model, calls the service layer to add the user,
// and returns a success or failure response to the client.

// @Summary Add a new user
// @Description Add a new user with the provided details.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.AddUserRequest true "User Details"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Security JWT
// @Router /api/v1/user [post]
func (h *userHandler) CreateUser(c *gin.Context) {
	var addUserRequest request.AddUserRequest

	// Bind JSON input to addUserRequest. Respond with a failure if the JSON is invalid.
	if err := c.ShouldBindJSON(&addUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Data:   fmt.Sprintf("Invalid request format: %v", err),
			Status: "fail",
		})
		return
	}

	// Call the service layer to add the user.
	user, err := h.service.CreateUser(addUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Data:   fmt.Sprintf("Failed to add user: %v", err),
			Status: "fail",
		})
		return
	}

	// Respond with a success message.
	c.JSON(http.StatusOK, response.APIResponse{
		Data:   user,
		Status: "success",
	})
}

// GetUser retrieves a specific user by their "id" value.
// It takes the "id" parameter from the request, calls the service layer, and returns the user data.
// @Summary Get a single user
// @Description Retrieve user details using the secondary key (sk).
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "Secondary Key (sk) of the user"
// @Success 200 {object} response.SingleUserResponse
// @Failure 400 {object} response.ErrorResponse
// @Security JWT
// @Router /api/v1/user/:id [get]
func (h *userHandler) ReadUser(c *gin.Context) {
	sk := c.Param("id") // Get the "id" parameter from the request.
	if sk == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "User ID is required",
		})
		return
	}

	user, err := h.service.ReadUser(sk)

	// Log and respond with an error if the service call fails.
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Message: fmt.Sprintf("Failed to retrieve user: %v", err),
		})
		return
	}

	// Respond with the user data and a success message.
	c.JSON(http.StatusOK, response.SingleUserResponse{
		Data:   user,
		Status: "success",
	})
}

// UpdateUser updates a user's data.
// It binds JSON input to a request model, calls the service layer to update the user,
// and returns the updated user or an error.
// @Summary Update a user's information
// @Description Update the details of a user identified by the secondary key (sk).
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "Secondary Key (sk) of the user"
// @Param request body request.UpdateUserRequest true "Update User Request Body"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Security JWT
// @Router /api/v1/user/:id [put]
func (h *userHandler) UpdateUser(c *gin.Context) {
	sk := c.Param("id") // Get the "id" parameter from the request.
	if sk == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Data:   "User ID is required",
			Status: "fail",
		})
		return
	}

	var updateUserRequest request.UpdateUserRequest

	// Bind JSON input to updateUserRequest. Respond with a failure if the JSON is invalid.
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Data:   fmt.Sprintf("Invalid request format: %v", err),
			Status: "fail",
		})
		return
	}

	// Call the service layer to update the user.
	user, err := h.service.UpdateUser(updateUserRequest, sk)

	// Log and respond with an error if the service call fails.
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Data:   fmt.Sprintf("Failed to update user: %v", err),
			Status: "fail",
		})
		return
	}

	// Respond with the updated user data and a success message.
	c.JSON(http.StatusOK, response.APIResponse{
		Data:   user,
		Status: "success",
	})
}

// DeleteUser deletes a user by their "id" value.
// It calls the service layer to delete the user and returns a success or failure response.
// @Summary Delete a user
// @Description Delete a user from the system using the secondary key (sk).
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "Secondary Key (sk) of the user"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Security JWT
// @Router /api/v1/user/:id [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {
	sk := c.Param("id") // Get the "id" parameter from the request.
	if sk == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Data:   "User ID is required",
			Status: "fail",
		})
		return
	}

	user, err := h.service.DeleteUser(sk)

	// Log and respond with an error if the service call fails.
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Data:   fmt.Sprintf("Failed to delete user: %v", err),
			Status: "fail",
		})
		return
	}

	// Respond with a success message.
	c.JSON(http.StatusOK, response.APIResponse{
		Data:   fmt.Sprintf("User deleted successfully (%s | %s)", user.Pk, user.Sk),
		Status: "success",
	})
}

// ListUsers retrieves the list of all users.
// It calls the service layer to get the list and returns it to the client.
// @Summary Get list of users
// @Description Retrieve the list of all users from the system.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} response.UsersListResponse
// @Failure 500 {object} response.ErrorResponse
// @Security JWT
// @Router /api/v1/user/list [get]
func (h *userHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers()

	// Log and respond with an error if the service call fails.
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.UsersListResponse{
			Data:   nil,
			Status: fmt.Sprintf("Failed to retrieve users: %v", err),
		})
		return
	}

	// Respond with the list of users and a success message.
	c.JSON(http.StatusOK, response.UsersListResponse{
		Data:   users,
		Status: "success",
	})
}
