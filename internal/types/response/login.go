package response

type LoginResponse struct {
	Token string `json:"token"`
}

type UserLoginRecord struct {
	Pk        string `json:"pk"`
	Sk        string `json:"sk"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
