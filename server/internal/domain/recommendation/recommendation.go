package recommendation

import "time"

type Season string

const (
	Spring Season = "Spring"
	Summer Season = "Summer"
	Autumn Season = "Autumn"
	Winter Season = "Winter"
)

type Outfit struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	UserID           uint      `gorm:"user_id"`
	Season           Season    `gorm:"season"`
	Occasion         string    `gorm:"occasion"`
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
