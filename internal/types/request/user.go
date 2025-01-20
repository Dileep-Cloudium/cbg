package request

type AddUserRequest struct {
	FirstName string `form:"first_name" json:"first_name"  binding:"required"`
	LastName  string `form:"last_name" json:"last_name"  binding:"required"`
	Email     string `form:"email" json:"email"  binding:"email"`
	Password  string `form:"password" json:"password"  binding:"required"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
