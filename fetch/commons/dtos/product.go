package dtos

import (
	"time"
)

type ProductList struct {
	Id         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	Price      string    `json:"price"`
	PriceIdr   string    `json:"price_idr"`
	Department string    `json:"department"`
	Product    string    `json:"product"`
}

type Recommended struct {
	Highest []ProductList `json:"highest"`
	Lowest  []ProductList `json:"lowest"`
}
