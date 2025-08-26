package Entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey" json:"id"`
	Phone string `gorm:"uniqueIndex" json:"phone"`
}
