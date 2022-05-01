package entities

type Permissions struct {
	ID          int      `json:"id" gorm:"primary_key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Roles       []*Roles `gorm:"many2many:roles_permissions;"`
}
