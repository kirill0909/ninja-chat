package models

type RegistrationRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegistrationResponse struct {
	Success string `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
}

type NonAuthHeaders struct {
	APIKey string `json:"Api-Key"`
}
