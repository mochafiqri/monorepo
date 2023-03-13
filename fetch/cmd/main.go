package main

import (
	"fetch/controllers"
	"fetch/domains"
	"fetch/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const PORT = ":8888"

func main() {
	e := echo.New()

	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})

	productRepo := repository.NewProductRepo()
	currencyRepo := repository.NewCurrencyRepo()
	productDomain := domains.NewProductDomain(productRepo, currencyRepo)
	productController := controllers.NewProductController(productDomain)
	productController.Routes(e)

	var err = e.Start(PORT)
	if err != nil {
		panic(err)
	}
}
