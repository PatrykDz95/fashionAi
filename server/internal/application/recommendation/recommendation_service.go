package recommendation

import (
	"fasion.ai/server/internal/domain/recommendation"
)

type RecommendationService struct {
	repo recommendation.Repository
}

func NewRecommendationService(repo recommendation.Repository) *RecommendationService {
	return &RecommendationService{repo: repo}
}

func (s *RecommendationService) GetRecommendations() ([]recommendation.Outfit, error) {
	return s.repo.GetRecommendations()
}

func (s *RecommendationService) GetRecommendationByID(id uint) (*recommendation.Outfit, error) {
	return s.repo.GetRecommendationByID(id)
}

func (s *RecommendationService) SaveRecommendation(outfit *recommendation.Outfit) error {
	return s.repo.SaveRecommendation(outfit)
}
