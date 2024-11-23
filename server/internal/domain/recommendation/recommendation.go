package recommendation

import "time"

type Outfit struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	UserID           uint      `gorm:"user_id"`
	Occasion         string    `gorm:"occasion"`
	Style            string    `gorm:"style"`
	RecommendedItems []Item    `gorm:"foreignKey:OutfitID"`
	DateCreated      time.Time `gorm:"date_created"`
}

type Item struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	OutfitID uint   `gorm:"not null"`
	Name     string `gorm:"name"`
	Brand    string `gorm:"brand"`
	Category string `gorm:"category"`
	Color    string `gorm:"color"`
	Material string `gorm:"material"`
}
