package user

type UserLoginRequest struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

type UserLoginResponse struct {
	Success      bool   `json:"success,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Error        string `json:"error,omitempty"`
}

type AuthData struct {
	UserID       int    `db:"id"`
	PasswordHash string `db:"password_hash"`
}
