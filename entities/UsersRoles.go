package entities

import "time"

type UsersRoles struct {
	UserID    int       `gorm:"primaryKey"`
	RoleID    int       `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
