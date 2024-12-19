package types

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name string
}
