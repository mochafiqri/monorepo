package repository

import (
	"encoding/json"
	"errors"
	"github.com/mochafiqri/monorepo/fetch/commons/constants"
	"github.com/mochafiqri/monorepo/fetch/commons/interfaces"
	"github.com/mochafiqri/monorepo/fetch/commons/utils"

	"net/url"
)

type CurrencyRepo struct {
	utilsHttp utils.HttpUtils
}

func NewCurrencyRepo() interfaces.CurrencyRepo {
	return &CurrencyRepo{utilsHttp: utils.GetHttpClient()}
}

func (r CurrencyRepo) GetIDR() (float64, error) {
	var (
		linkUrl = "https://api.freecurrencyapi.com/v1/latest"
	)

	var urlParam = url.Values{}
	urlParam.Add("base_currency", "USD")
	urlParam.Add("currencies", "IDR")
	urlParam.Add("apikey", constants.CurrencyApiKey)

	var header = map[string]string{
		"apikey": constants.CurrencyApiKey,
	}

	resHttp, err := r.utilsHttp.DoRequest("GET", linkUrl, "", nil, header)
	if err != nil {
		return 0, err
	}

	if resHttp.Code != 200 {
		return 0, errors.New("currency might be down")
	}

	var body struct {
		Data struct {
			IDR float64 `json:"IDR"`
		} `json:"data"`
	}

	err = json.Unmarshal(resHttp.Body, &body)
	if err != nil {
		return 0, errors.New("source product might be down")
	}

	return body.Data.IDR, nil
}
