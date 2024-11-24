package recommendation

import (
	"fasion.ai/server/internal/domain/recommendation"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) recommendation.Repository {
	return &repository{db: db}
}

func (r *repository) GetRecommendations() ([]recommendation.Outfit, error) {
	var recommendations []recommendation.Outfit
	if err := r.db.Preload("RecommendedItems").Find(&recommendations).Error; err != nil {
		return nil, err
	}
	return recommendations, nil
}

func (r *repository) GetRecommendationByID(id uint) (*recommendation.Outfit, error) {
	var recommendation recommendation.Outfit
	if err := r.db.Preload("RecommendedItems").First(&recommendation, id).Error; err != nil {
		return nil, err
	}
	return &recommendation, nil
}

func (r *repository) SaveRecommendation(outfit *recommendation.Outfit) error {
	if err := r.db.Create(outfit).Error; err != nil {
		return err
	}
	return nil
}
