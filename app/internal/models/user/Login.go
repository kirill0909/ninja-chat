package user

type UserLoginRequest struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

type UserLoginResponse struct {
	Success     bool   `json:"success,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	Error       string `json:"error,omitempty"`
	Code        int    `json:"code,omitempty"`
}

type AuthData struct {
	UserID       int    `db:"id"`
	PasswordHash string `db:"password_hash"`
}

type TokenData struct {
	AccessToken string
}

type UserSession struct {
	UserID      int    `json:"userID"`
	AccessToken string `json:"accessToken"`
	ExpiredAt   int    `json:"expiredAt"`
}
