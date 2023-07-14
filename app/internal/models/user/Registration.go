package models

type RegistrationRequest struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}
