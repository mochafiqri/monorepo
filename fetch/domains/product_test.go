package domains

import (
	"github.com/golang/mock/gomock"
	"github.com/mochafiqri/monorepo/fetch/commons/constants"
	"github.com/mochafiqri/monorepo/fetch/commons/dtos"
	"github.com/mochafiqri/monorepo/fetch/commons/entities"
	"github.com/mochafiqri/monorepo/fetch/commons/interfaces"
	mock_interfaces "github.com/mochafiqri/monorepo/fetch/mock"
	"net/http"
	"time"

	"reflect"
	"testing"
)

func TestProductDomain_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var productRepo = mock_interfaces.NewMockProductRepo(ctrl)
	var currencyRepo = mock_interfaces.NewMockCurrencyRepo(ctrl)

	type fields struct {
		productRepo  interfaces.ProductRepo
		currencyRepo interfaces.CurrencyRepo
	}
	var tests = []struct {
		name   string
		fields fields
		want   dtos.StandardResponseReq
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				productRepo:  productRepo,
				currencyRepo: currencyRepo,
			},
			want: dtos.StandardResponseReq{Code: http.StatusOK, Data: []dtos.ProductList{{
				Id:         "1",
				CreatedAt:  time.Time{},
				Price:      "1.00",
				PriceIdr:   "15000",
				Department: "penerangan",
				Product:    "A",
			}}, Meta: map[string]interface{}{
				"count": 1,
			}},
			mock: func() {
				data := []entities.Product{{
					Id:         "1",
					CreatedAt:  time.Time{},
					Price:      "1.00",
					Department: "penerangan",
					Product:    "A",
				}}
				var idr float64 = 15000
				productRepo.EXPECT().GetProduct().Return(data, nil)
				currencyRepo.EXPECT().GetIDR().Return(idr, nil)
			},
		},
		{
			name: "not found",
			fields: fields{
				productRepo:  productRepo,
				currencyRepo: currencyRepo,
			},
			want: dtos.StandardResponseReq{Code: http.StatusNotFound, Message: constants.MsgDataNotFound},
			mock: func() {
				var idr float64 = 15000
				productRepo.EXPECT().GetProduct().Return(nil, nil)
				currencyRepo.EXPECT().GetIDR().Return(idr, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			d := ProductDomain{
				productRepo:  tt.fields.productRepo,
				currencyRepo: tt.fields.currencyRepo,
			}

			if got := d.GetProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductDomain_GetProductRecommended(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var productRepo = mock_interfaces.NewMockProductRepo(ctrl)
	var currencyRepo = mock_interfaces.NewMockCurrencyRepo(ctrl)

	type fields struct {
		productRepo  interfaces.ProductRepo
		currencyRepo interfaces.CurrencyRepo
	}

	var idr float64 = 15000

	var tests = []struct {
		name   string
		fields fields
		want   dtos.StandardResponseReq
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				productRepo:  productRepo,
				currencyRepo: currencyRepo,
			},
			want: dtos.StandardResponseReq{Code: http.StatusOK, Data: expectRecommend, Meta: map[string]interface{}{
				"idr": idr,
			}},
			mock: func() {
				data := dataRecommend
				productRepo.EXPECT().GetProduct().Return(data, nil)
				currencyRepo.EXPECT().GetIDR().Return(idr, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			d := ProductDomain{
				productRepo:  tt.fields.productRepo,
				currencyRepo: tt.fields.currencyRepo,
			}

			if got := d.GetProductRecommended(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductRecommended() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helperProductToProductList(t *testing.T) {
	type args struct {
		products []entities.Product
		idr      float64
	}
	tests := []struct {
		name string
		args args
		want []dtos.ProductList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helperProductToProductList(tt.args.products, tt.args.idr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helperProductToProductList() = %v, want %v", got, tt.want)
			}
		})
	}
}

var dataRecommend = []entities.Product{
	{
		Id:         "1",
		CreatedAt:  time.Time{},
		Price:      "1.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "2",
		CreatedAt:  time.Time{},
		Price:      "2.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "3",
		CreatedAt:  time.Time{},
		Price:      "3.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "4",
		CreatedAt:  time.Time{},
		Price:      "4.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "5",
		CreatedAt:  time.Time{},
		Price:      "5.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "6",
		CreatedAt:  time.Time{},
		Price:      "6.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "7",
		CreatedAt:  time.Time{},
		Price:      "7.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "8",
		CreatedAt:  time.Time{},
		Price:      "8.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "9",
		CreatedAt:  time.Time{},
		Price:      "9.00",
		Department: "penerangan",
		Product:    "A",
	}, {
		Id:         "10",
		CreatedAt:  time.Time{},
		Price:      "10.00",
		Department: "penerangan",
		Product:    "A",
	}}

var expectRecommend = dtos.Recommended{
	Highest: []dtos.ProductList{
		{
			Id:         "10",
			CreatedAt:  time.Time{},
			Price:      "10.00",
			PriceIdr:   "150000",
			Department: "penerangan",
			Product:    "A",
		}, {
			Id:         "9",
			CreatedAt:  time.Time{},
			Price:      "9.00",
			PriceIdr:   "135000",
			Department: "penerangan",
			Product:    "A",
		}, {
			Id:         "8",
			CreatedAt:  time.Time{},
			Price:      "8.00",
			PriceIdr:   "120000",
			Department: "penerangan",
			Product:    "A",
		},
		{
			Id:         "7",
			CreatedAt:  time.Time{},
			Price:      "7.00",
			PriceIdr:   "105000",
			Department: "penerangan",
			Product:    "A",
		},
		{
			Id:         "6",
			CreatedAt:  time.Time{},
			Price:      "6.00",
			PriceIdr:   "90000",
			Department: "penerangan",
			Product:    "A",
		},
	},
	Lowest: []dtos.ProductList{
		{
			Id:         "1",
			CreatedAt:  time.Time{},
			Price:      "1.00",
			PriceIdr:   "15000",
			Department: "penerangan",
			Product:    "A",
		}, {
			Id:         "2",
			CreatedAt:  time.Time{},
			Price:      "2.00",
			PriceIdr:   "30000",
			Department: "penerangan",
			Product:    "A",
		}, {
			Id:         "3",
			CreatedAt:  time.Time{},
			Price:      "3.00",
			PriceIdr:   "45000",
			Department: "penerangan",
			Product:    "A",
		}, {
			Id:         "4",
			CreatedAt:  time.Time{},
			Price:      "4.00",
			PriceIdr:   "60000",
			Department: "penerangan",
			Product:    "A",
		}, {
			Id:         "5",
			CreatedAt:  time.Time{},
			Price:      "5.00",
			PriceIdr:   "75000",
			Department: "penerangan",
			Product:    "A",
		}},
}
