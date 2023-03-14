package interfaces

import (
	"github.com/mochafiqri/monorepo/fetch/commons/dtos"
	"github.com/mochafiqri/monorepo/fetch/commons/entities"
)

type ProductRepo interface {
	GetProduct() ([]entities.Product, error)
}

type ProductDomain interface {
	GetProduct() dtos.StandardResponseReq
	GetProductRecommended() dtos.StandardResponseReq
}
