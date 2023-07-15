package user

type UserLoginRequest struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

type UserLoginResponse struct {
	Success      bool   `json:"success"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Error        string `json:"error,omitempty"`
}
