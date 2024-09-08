package preferences

type Preference struct {
	ID             uint     `gorm:"primaryKey;autoIncrement"`
	UserID         uint     `gorm:"user_id"`
	Style          string   `gorm:"style"`
	ColorPalette   []string `gorm:"color_palette"`
	Fit            string   `gorm:"fit"`
	FavoriteBrands []string `gorm:"favorite_brands"`
	BudgetRange    string   `gorm:"budget_range"`
}
