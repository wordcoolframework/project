package Contracts

import "time"

type IOtpRepsitory interface {
	Create(phone, otp string, duration time.Duration) error

	Verify(phone, otp string) (bool, error)

	CountOtpRequestsInLastDuration(phone string, duration time.Duration) (int64, error)
}
