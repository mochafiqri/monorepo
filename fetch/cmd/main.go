package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mochafiqri/monorepo/fetch/controllers"
	_ "github.com/mochafiqri/monorepo/fetch/docs"
	"github.com/mochafiqri/monorepo/fetch/domains"
	"github.com/mochafiqri/monorepo/fetch/repository"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const PORT = ":8888"

// @title Swagger Fetch App
// @version 1.0
// @description This is a featch server for JDS Technical.
// @termsOfService http://swagger.io/terms/

// @contact.name Mochamad Fiqri
// @contact.url https://github.com/mochafiqri
// @contact.email mocha.fiqri@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
