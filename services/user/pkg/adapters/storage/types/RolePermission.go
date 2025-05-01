package types

type RolePermission struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"index"`
	Role       string `gorm:"index;size:255"`
	Permission string `gorm:"size:255"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
