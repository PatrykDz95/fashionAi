package auth

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"not null"`
	Email        string    `gorm:"not null; unique"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"not null"`
	IsActive     bool      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null; autoCreateTime"`
	UpdatedAt    time.Time `gorm:"not null; autoUpdateTime"`
}
