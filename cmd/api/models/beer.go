//models.beer

package models

// BeerItem model.
type BeerItem struct {
	ModelBase
	Name     string  `json:"name" gorm:"index" binding:"required"`
	Brewery  string  `json:"brewery" binding:"required"`
	Country  string  `json:"country" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type BeerBox struct {
	PriceTotal float64 `json:"price_total"`
}
