package models

type RegistrationRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type NonAuthHeaders struct {
	APIKey string `json:"Api-Key"`
}
