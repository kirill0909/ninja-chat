package user

import "strings"

type RegistrationRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegistrationResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (r *RegistrationRequest) Validate() bool {
	if strings.TrimSpace(r.Login) == "" ||
		strings.TrimSpace(r.Password) == "" {
		return false
	}
	return true
}
