package conn

type RespData struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}
