package response

type APIResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Status string      `json:"status"`
}
