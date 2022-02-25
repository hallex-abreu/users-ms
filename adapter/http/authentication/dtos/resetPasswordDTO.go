package dtos

type ResetPasswordDTO struct {
	Token    string `json:"token" db:"password_reset_token"`
	Password string `json:"password" db:"password"`
}
