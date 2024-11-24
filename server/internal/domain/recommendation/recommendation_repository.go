package recommendation

type Repository interface {
	GetRecommendations() ([]Outfit, error)
	GetRecommendationByID(id uint) (*Outfit, error)
	SaveRecommendation(outfit *Outfit) error
}
