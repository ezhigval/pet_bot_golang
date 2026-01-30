package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"zelential.com/user-service/models"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Создать пользователя
func (r *UserRepository) CreateUser(user *models.User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := r.DB.NamedExec(`
		INSERT INTO users (id,email,password_hash,device_ids,two_fa_enabled,created_at,updated_at,trial_expires)
		VALUES (:id,:email,:password_hash,:device_ids,:two_fa_enabled,:created_at,:updated_at,:trial_expires)
	`, user)
	return err
}

// Найти по email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

// Проверка устройства
func (r *UserRepository) IsDeviceUsed(deviceID string) (bool, error) {
	var count int
	err := r.DB.Get(&count, "SELECT COUNT(*) FROM users WHERE $1=ANY(device_ids)", deviceID)
	return count > 0, err
}

// Добавление устройства
func (r *UserRepository) AddDevice(userID, deviceID string) error {
	_, err := r.DB.Exec(`
		UPDATE users SET device_ids = array_append(device_ids, $2), updated_at=$3 WHERE id=$1
	`, userID, deviceID, time.Now())
	return err
}
