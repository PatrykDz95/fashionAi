package recommendation

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewRecommendationService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) GetRecommendations() ([]Outfit, error) {
	var recommendations []Outfit
	if err := s.DB.Find(&recommendations).Error; err != nil {
		return nil, err
	}
	return recommendations, nil
}

func (s *Service) GetRecommendationByID(id string) (*Outfit, error) {
	var recommendation Outfit
	if err := s.DB.First(&recommendation, id).Error; err != nil {
		return nil, err
	}
	return &recommendation, nil
}
