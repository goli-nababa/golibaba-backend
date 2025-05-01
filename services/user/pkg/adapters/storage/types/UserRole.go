package types

type UserRole struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"index"`
	Role   string `gorm:"size:255"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
