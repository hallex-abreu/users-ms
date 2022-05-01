package entities

type Roles struct {
	ID          int            `json:"id" gorm:"primary_key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Users       []*Users       `gorm:"many2many:users_roles;"`
	Permissions []*Permissions `gorm:"many2many:roles_permissions;"`
}
