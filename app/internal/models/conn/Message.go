package conn

type SendMessageRequest struct {
	UserIDFrom int    `json:"userIDFrom" validate:"required"`
	UserIDTo   int    `json:"userIDTo" validate:"required"`
	Message    string `json:"message" validate:"required"`
}

type SendMessageResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}
