package dtos

type RecoverPasswordDTO struct {
	Email string `json:"email" db:"email"`
}
