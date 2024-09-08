package recommendation

import "time"

type Outfit struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	UserID           uint      `gorm:"user_id"`
	Occasion         string    `gorm:"occasion"`
	Style            string    `gorm:"style"`
	RecommendedItems []Item    `gorm:"recommended_items"`
	DateCreated      time.Time `gorm:"date_created"`
}

type Item struct {
	ID                     uint    `gorm:"primaryKey;autoIncrement"`
	OutfitRecommendationID uint    `gorm:"not null"`
	Name                   string  `gorm:"name"`
	Brand                  string  `gorm:"brand"`
	Category               string  `gorm:"category"`
	Size                   string  `gorm:"size"`
	Price                  float64 `gorm:"price"`
	Image                  string  `gorm:"image"`
	Color                  string  `gorm:"color"`
	Material               string  `gorm:"material"`
	Retailer               string  `gorm:"retailer"`
	Rating                 float64 `gorm:"rating"`
	AffiliateLink          string  `gorm:"affiliate_link"`
}
