package interfaces

import (
	"fetch/commons/dtos"
	"fetch/commons/entities"
)

type ProductRepo interface {
	GetProduct() ([]entities.Product, error)
}

type ProductDomain interface {
	GetProduct() dtos.StandardResponseReq
	GetProductRecommended() dtos.StandardResponseReq
}
