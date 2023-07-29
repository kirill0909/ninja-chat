package conn

type SendMessageRequest struct {
	RecipientID int    `json:"recipientID" validate:"required"`
	Message     string `json:"message" validate:"required"`
	UserID      int
}

type SendMessageResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}
