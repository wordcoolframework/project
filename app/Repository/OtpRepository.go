package Repository

import (
	"gorm.io/gorm"
	"projectUserManagement/app/Domain/Contracts"
	"projectUserManagement/app/Domain/Entity"
	"time"
)

type OtpRepository struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) Contracts.IOtpRepsitory {
	return &OtpRepository{db: db}
}

func (r *OtpRepository) Create(phone, otp string, duration time.Duration) error {
	return r.db.Create(&Entity.Otp{
		Phone:     phone,
		OTP:       otp,
		ExpiresAt: time.Now().Add(duration),
	}).Error
}

func (r *OtpRepository) Verify(phone, otp string) (bool, error) {
	var record Entity.Otp
	err := r.db.Where("phone = ? AND otp = ?", phone, otp).
		Order("created_at DESC").
		First(&record).Error
	if err != nil {
		return false, err
	}

	if time.Now().After(record.ExpiresAt) {
		return false, nil
	}

	r.db.Delete(&record)
	return true, nil
}

func (r *OtpRepository) CountOtpRequestsInLastDuration(phone string, duration time.Duration) (int64, error) {
	var count int64
	thresholdTime := time.Now().Add(-duration)
	err := r.db.Model(&Entity.Otp{}).
		Where("phone = ? AND created_at >= ?", phone, thresholdTime).
		Count(&count).Error
	return count, err
}
