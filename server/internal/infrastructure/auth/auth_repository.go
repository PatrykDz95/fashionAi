package auth

import (
	"fasion.ai/server/internal/domain/auth"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *auth.User) error
	GetUserByUsername(username string) (*auth.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user *auth.User) error {
	return r.db.Create(user).Error
}

func (r *repository) GetUserByUsername(username string) (*auth.User, error) {
	var user auth.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
