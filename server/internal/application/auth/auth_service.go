package auth

import (
	"fasion.ai/server/internal/domain/auth"
	infraAuth "fasion.ai/server/internal/infrastructure/auth"
)

type AuthService struct {
	repo infraAuth.Repository
}

func NewAuthService(repo infraAuth.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(username string) (string, error) {
	return auth.GenerateToken(username)
}

func (s *AuthService) CreateUser(user *auth.User) error {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserByUsername(username string) (*auth.User, error) {
	return s.repo.GetUserByUsername(username)
}
