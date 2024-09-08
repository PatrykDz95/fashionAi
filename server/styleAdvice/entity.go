package styleAdvice

import "fasion.ai/server/recommendation"

type StyleAdvice struct {
	ID           uint                  `gorm:"primaryKey;autoIncrement"`
	UserID       uint                  `gorm:"user_id"`
	Prompt       string                `gorm:"prompt"`
	Response     string                `gorm:"response"`
	Outfit       recommendation.Outfit `gorm:"not null"`
	DateCreated  string                `gorm:"date_created"`
	UserFeedback string                `gorm:"user_feedback"`
}
