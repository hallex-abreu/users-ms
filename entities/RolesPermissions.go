package entities

import "time"

type RolesPermissions struct {
	RolesID       int       `gorm:"primaryKey"`
	PermissionsID int       `gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
