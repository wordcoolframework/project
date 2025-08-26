package Service

import (
	"errors"
	"fmt"
	"math/rand"
	"projectUserManagement/app/Domain/Entity"
	"projectUserManagement/app/Repository"
	"projectUserManagement/app/jwt"
	"time"
)

type AuthService struct {
	userRepo  *Repository.UserRepository
	otpRepo   *Repository.OtpRepository
	jwtSecret string
}

func generateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(900000)+100000)
}

func NewAuthService(userRepo *Repository.UserRepository, otpRepo *Repository.OtpRepository, secret string) *AuthService {
	return &AuthService{userRepo, otpRepo, secret}
}

func (s *AuthService) RequestOTP(phone string) (string, error) {
	count, _ := s.otpRepo.CountOtpRequestsInLastDuration(phone, 10*time.Minute)
	if count >= 3 {
		return "", errors.New("rate limitation")
	}

	otp := generateOTP()
	err := s.otpRepo.Create(phone, otp, 2*time.Minute)
	if err != nil {
		return "", err
	}

	println("otp for", phone, "is", otp)
	return otp, nil
}

func (s *AuthService) VerifyOTP(phone, otp string) (string, *Entity.User, error) {
	ok, _ := s.otpRepo.Verify(phone, otp)

	if !ok {
		return "", nil, errors.New("invalid or expired otp code")
	}

	user, err := s.userRepo.GetByPhone(phone)

	if err != nil {

		user = &Entity.User{
			Phone: phone,
		}
		err := s.userRepo.Create(user)

		if err != nil {
			return "", nil, err
		}
	}

	token, _ := jwt.GenerateToken(user.ID, phone, s.jwtSecret)
	return token, user, nil
}
