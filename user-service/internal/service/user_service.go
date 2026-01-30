package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"zelential.com/user-service/models"
	"zelential.com/user-service/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Регистрация
func (s *UserService) Register(email, password, deviceID string) (*models.User, error) {
	existing, _ := s.Repo.GetByEmail(email)
	deviceUsed, _ := s.Repo.IsDeviceUsed(deviceID)

	if deviceUsed {
		// Ограничение пробного периода при повторной регистрации
		if existing != nil {
			existing.TrialExpires = time.Now().Add(7 * 24 * time.Hour) // сокращенный пробный период
			return existing, errors.New("device already used, trial shortened")
		} else {
			return nil, errors.New("this device already has a registered account")
		}
	}

	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &models.User{
		Email:        email,
		PasswordHash: string(hash),
		DeviceIDs:    []string{deviceID},
		TwoFAEnabled: false,
		TrialExpires: time.Now().Add(14 * 24 * time.Hour),
	}

	err := s.Repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Проверка пароля
func (s *UserService) VerifyPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}
