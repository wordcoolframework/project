package Entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone string `gorm:"uniqueIndex" json:"phone"`
}
