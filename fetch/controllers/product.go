package controllers

import (
	"fetch/commons/interfaces"
	"fetch/commons/utils"
	"fetch/middleware"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productDomain interfaces.ProductDomain
}

func NewProductController(productDomain interfaces.ProductDomain) *ProductController {
	return &ProductController{
		productDomain: productDomain,
	}
}

func (c *ProductController) GetProduct(e echo.Context) error {
	var resp = c.productDomain.GetProduct()
	return utils.WriteResponse(e, resp)
}

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
