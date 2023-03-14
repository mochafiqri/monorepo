package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/monorepo/fetch/commons/interfaces"
	"github.com/mochafiqri/monorepo/fetch/commons/utils"
	"github.com/mochafiqri/monorepo/fetch/middleware"
)

type ProductController struct {
	productDomain interfaces.ProductDomain
}

func NewProductController(productDomain interfaces.ProductDomain) *ProductController {
	return &ProductController{
		productDomain: productDomain,
	}
}

// GetProduct godoc
// @Summary Get Product From Mock Api
// @Tags Product
// @Accept  json
// @Produce  json
// @Param authorization header string true "With the Bearer started"
// @Success 200 {object} dtos.StandardResponse{data=[]dtos.ProductList}
// @Failure 401 {object} dtos.StandardResponse{}
// @Failure 400 {object} dtos.StandardResponse{}
// @Failure 500 {object} dtos.StandardResponse{}
// @Router /api/v1/products [get]
func (c *ProductController) GetProduct(e echo.Context) error {
	var resp = c.productDomain.GetProduct()
	return utils.WriteResponse(e, resp)
}

// GetProductRecommended godoc
// @Summary Get Product Recommend (Highest-Lowest) by Price
// @Tags Product
// @Accept  json
// @Produce  json
// @Param authorization header string true "With the Bearer started"
// @Success 200 {object} dtos.StandardResponse{data=[]dtos.ProductList}
// @Failure 401 {object} dtos.StandardResponse{}
// @Failure 400 {object} dtos.StandardResponse{}
// @Failure 500 {object} dtos.StandardResponse{}
// @Router /api/v1/products/recommended [get]
func (c *ProductController) GetProductRecommended(e echo.Context) error {
	var resp = c.productDomain.GetProductRecommended()
	return utils.WriteResponse(e, resp)
}

func (c *ProductController) Routes(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	v1.Use(middleware.Auth)
	v1.GET("/products", c.GetProduct)
	v1.GET("/products/recommended", c.GetProductRecommended)
}
