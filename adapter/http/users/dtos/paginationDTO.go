package dtos

type PaginationDTO struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
