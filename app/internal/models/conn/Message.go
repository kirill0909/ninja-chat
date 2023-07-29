package conn

type SaveMessageRequest struct {
	RecipientID int    `json:"recipientID" validate:"required"`
	Message     string `json:"message" validate:"required"`
	MessageID   int
	UserID      int
}

type SaveMessageResponse struct {
	MessageID int    `json:"messageID,omitempty"`
	Success   bool   `json:"success,omitempty"`
	Error     string `json:"error,omitempty"`
	Code      int    `json:"code,omitempty"`
}
