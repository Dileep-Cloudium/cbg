package request

type LoginRequest struct {
	Email    string `form:"email" json:"email"  binding:"email"`
	Password string `form:"password" json:"password"  binding:"required"`
}
