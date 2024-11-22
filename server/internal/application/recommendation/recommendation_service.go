package recommendation

import (
	domain "fasion.ai/server/internal/domain/recommendation"
	"fasion.ai/server/internal/infrastructure/recommendation"
)

type RecommendationService struct {
	repo recommendation.Repository
}

func NewRecommendationService(repo recommendation.Repository) *RecommendationService {
	return &RecommendationService{repo: repo}
}

func (s *RecommendationService) GetRecommendations() ([]domain.Outfit, error) {
	return s.repo.GetRecommendations()
}

func (s *RecommendationService) GetRecommendationByID(id uint) (*domain.Outfit, error) {
	return s.repo.GetRecommendationByID(id)
}

func (s *RecommendationService) SaveRecommendation(outfit *domain.Outfit) error {
	return s.repo.SaveRecommendation(outfit)
}
