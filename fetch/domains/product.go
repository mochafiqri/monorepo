package domains

import (
	"github.com/mochafiqri/monorepo/fetch/commons/constants"
	"github.com/mochafiqri/monorepo/fetch/commons/dtos"
	"github.com/mochafiqri/monorepo/fetch/commons/entities"
	"github.com/mochafiqri/monorepo/fetch/commons/interfaces"
	"github.com/sirupsen/logrus"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

type ProductDomain struct {
	productRepo  interfaces.ProductRepo
	currencyRepo interfaces.CurrencyRepo
}

func NewProductDomain(productRepo interfaces.ProductRepo, currencyRepo interfaces.CurrencyRepo) interfaces.ProductDomain {
	return &ProductDomain{
		productRepo:  productRepo,
		currencyRepo: currencyRepo,
	}
}

func (d ProductDomain) GetProduct() dtos.StandardResponseReq {
	var (
		productResp         = make([]dtos.ProductList, 0)
		idr         float64 = 0
	)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		tmpIdr, err := d.currencyRepo.GetIDR()
		if err != nil {
			logrus.Error("[GetProduct][GetIDR] err ", err.Error())
		}
		idr = tmpIdr
		wg.Done()
	}()

	product, err := d.productRepo.GetProduct()
	if err != nil {
		return dtos.StandardResponseReq{Code: http.StatusInternalServerError, Error: err}
	}

	wg.Wait()

	if len(product) == 0 {
		return dtos.StandardResponseReq{Code: http.StatusNotFound, Message: constants.MsgDataNotFound}

	}

	productResp = helperProductToProductList(product, idr)

	var meta = map[string]interface{}{
		"count": len(product),
	}

	return dtos.StandardResponseReq{Code: http.StatusOK, Data: productResp, Meta: meta}
}

func (d ProductDomain) GetProductRecommended() dtos.StandardResponseReq {
	var (
		highest         = make([]entities.Product, 0)
		lowest          = make([]entities.Product, 0)
		resp            = dtos.Recommended{}
		idr     float64 = 0
	)
	product, err := d.productRepo.GetProduct()
	if err != nil {
		return dtos.StandardResponseReq{Code: http.StatusInternalServerError, Error: err}
	}

	if len(product) == 0 {
		return dtos.StandardResponseReq{Code: http.StatusNotFound}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		tmpIdr, err := d.currencyRepo.GetIDR()
		if err != nil {
			logrus.Error("[GetProduct][GetIDR] err ", err.Error())
		}
		idr = tmpIdr
		wg.Done()
	}()

	for i, v := range product {
		product[i].PriceFloat, _ = strconv.ParseFloat(v.Price, 64)
	}

	sort.Slice(product, func(i, j int) bool {
		return product[i].PriceFloat < product[j].PriceFloat
	})

	if len(product) < 5 {
	} else {
		for i := 0; i < 5; i++ {
			lowest = append(lowest, product[i])
			highest = append(highest, product[len(product)-(i+1)])
		}
	}
	wg.Wait()
	resp.Highest = helperProductToProductList(highest, idr)
	resp.Lowest = helperProductToProductList(lowest, idr)
	var meta = map[string]interface{}{
		"idr": idr,
	}
	return dtos.StandardResponseReq{Data: resp, Code: http.StatusOK, Meta: meta}
}

func helperProductToProductList(products []entities.Product, idr float64) []dtos.ProductList {
	var data = make([]dtos.ProductList, 0)

	for _, v := range products {
		var priceIdr float64 = 0
		if idr != 0 {
			tmpPrice, err := strconv.ParseFloat(v.Price, 64)
			if err != nil {
				logrus.Error("[GetProduct][ParseFloat] err ", err.Error())
			} else {
				priceIdr = tmpPrice * idr
			}
		}
		data = append(data, dtos.ProductList{
			Id:         v.Id,
			CreatedAt:  v.CreatedAt,
			Price:      v.Price,
			PriceIdr:   strconv.Itoa(int(priceIdr)),
			Department: v.Department,
			Product:    v.Product,
		})
	}

	return data

}
