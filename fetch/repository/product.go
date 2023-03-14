package repository

import (
	"encoding/json"
	"errors"
	"github.com/mochafiqri/monorepo/fetch/commons/entities"
	"github.com/mochafiqri/monorepo/fetch/commons/interfaces"
	"github.com/mochafiqri/monorepo/fetch/commons/utils"
)

var urlProduct = "https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product"

type ProductRepo struct {
	utilsHttp utils.HttpUtils
}

func NewProductRepo() interfaces.ProductRepo {
	return &ProductRepo{utilsHttp: utils.GetHttpClient()}
}

func (r ProductRepo) GetProduct() ([]entities.Product, error) {
	var (
		url = urlProduct
		res []entities.Product
	)

	resHttp, err := r.utilsHttp.DoRequest("GET", url, "", nil, nil)
	if err != nil {
		return nil, err
	}

	if resHttp.Code != 200 {
		return nil, errors.New("source product might be down")
	}

	err = json.Unmarshal(resHttp.Body, &res)
	if err != nil {
		return nil, errors.New("source product might be down")
	}

	return res, nil
}
