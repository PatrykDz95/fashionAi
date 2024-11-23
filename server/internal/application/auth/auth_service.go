package auth

import (
	"os"
	"time"

	domain_auth "fasion.ai/server/internal/domain/auth"
	infra_auth "fasion.ai/server/internal/infrastructure/auth"
	"github.com/golang-jwt/jwt/v5"
)

var appName = os.Getenv("APP_NAME")
var jwtSecretKey = os.Getenv("JWT_SECRET")
var jwtSecret = []byte(jwtSecretKey)

type AuthService struct {
	repo infra_auth.Repository
}

func NewAuthService(repo infra_auth.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"iat":      time.Now().Unix(),
		"iss":      appName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) CreateUser(user *domain_auth.User) error {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserByUsername(username string) (*domain_auth.User, error) {
	return s.repo.GetUserByUsername(username)
}
