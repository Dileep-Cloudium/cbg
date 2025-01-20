package response

type SingleUserResponse struct {
	Data   CommonResponse `json:"data"`
	Status string         `json:"status"`
}

type AddUserResponse struct {
	Data   string `json:"data"`
	Status string `json:"status"`
}

type CommonResponse struct {
	Pk        string `json:"pk"`
	Sk        string `json:"sk"`
	ItemBody  string `json:"item_body"`
	CreatedAt string `json:"created_at"`
}

type UsersListResponse struct {
	Data   []CommonResponse `json:"data"`
	Status string           `json:"status"`
}
