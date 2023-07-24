package user

type LogoutResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code"`
}
