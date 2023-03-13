package entities

import "time"

type Product struct {
	Id         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	Price      string    `json:"price"`
	PriceFloat float64   `json:"-"`
	Department string    `json:"department"`
	Product    string    `json:"product"`
}
