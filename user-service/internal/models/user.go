package models

import (
	"time"

	"github.com/google/uuid"
)

// User — структура для хранения информации о пользователе
type User struct {
	ID           uuid.UUID `db:"id"`             // уникальный идентификатор
	Email        string    `db:"email"`          // email пользователя
	Phone        string    `db:"phone"`          // телефон
	PasswordHash string    `db:"password_hash"`  // хеш пароля
	CreatedAt    time.Time `db:"created_at"`     // дата создания аккаунта
	UpdatedAt    time.Time `db:"updated_at"`     // дата последнего обновления
	LastLogin    time.Time `db:"last_login"`     // дата последнего входа
	IsActive     bool      `db:"is_active"`      // активный аккаунт
	TrialExpires time.Time `db:"trial_expires"`  // дата окончания пробного периода
	TwoFAEnabled bool      `db:"two_fa_enabled"` // включена ли двухфакторная аутентификация
	DeviceIDs    []string  `db:"device_ids"`     // список ID привязанных устройств
	IPAddresses  []string  `db:"ip_addresses"`   // список последних IP
	Suspicious   bool      `db:"suspicious"`     // флаг подозрительной активности
}
