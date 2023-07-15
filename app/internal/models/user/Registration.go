package user

type RegistrationRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegistrationResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
}

type NonAuthHeaders struct {
	APIKey string `json:"Api-Key"`
}
