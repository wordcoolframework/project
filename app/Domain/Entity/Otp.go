package Entity

import (
	"gorm.io/gorm"
	"time"
)

type OTPRequest struct {
	gorm.Model
	Phone     string `gorm:"index"`
	OTP       string
	ExpiresAt time.Time
}
