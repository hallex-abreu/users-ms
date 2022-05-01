package entities

import "time"

type Users struct {
	ID                   int       `json:"id" gorm:"primary_key"`
	Name                 string    `json:"name"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	PasswordResetToken   string    `json:"password_reset_token" gorm:"default:null"`
	PasswordResetExpires int64     `json:"password_reset_expires" gorm:"default:null"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Roles                []*Roles  `gorm:"many2many:users_roles;"`
}
