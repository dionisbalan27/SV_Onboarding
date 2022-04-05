package product_usecase

import (
	"backend-api/helpers"
	"backend-api/models/product/dto"
	"backend-api/models/product/entity"
	"errors"

	"github.com/stretchr/testify/mock"
)

type ProductUsecaseMockInterface interface {
	GetAllProducts() dto.Response
	GetProductById(id string) dto.Response
	UpdateProductData(id string) dto.Response
	CreateNewProduct(product entity.Product) dto.Response
	DeleteProductById(id string) dto.Response
}

type ProductUsecaseMock struct {
	Mock mock.Mock
}

func (product *ProductUsecaseMock) GetAllProducts() dto.Response {
	arguments := product.Mock.Called()
	if arguments.Get(0) == nil {
		return helpers.ResponseError("Data not found", errors.New("No record"), 404)
	}
	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)

}

func (product *ProductUsecaseMock) GetProductById(id string) dto.Response {
	arguments := product.Mock.Called(id)
	if arguments.Get(0) == nil {
		return helpers.ResponseError("Data not found", errors.New("No record"), 404)
	}
	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}

func (product *ProductUsecaseMock) UpdateProductData(id string) dto.Response {
	arguments := product.Mock.Called(id)

	if arguments.Get(0) == nil {
		return helpers.ResponseError("Data not found", errors.New("Failed to update product"), 404)
	}
	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}

func (product *ProductUsecaseMock) DeleteProductById(id string) dto.Response {
	arguments := product.Mock.Called(id)
	if arguments.Get(0) == nil {
		return helpers.ResponseError("Data not found", errors.New("Failed to delete product"), 404)
	}

	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}

func (product *ProductUsecaseMock) CreateNewProduct(productData dto.Product) dto.Response {
	arguments := product.Mock.Called()
	if arguments.Get(0) == nil {
		return helpers.ResponseError("Internal server error", errors.New("Failed to create Product"), 500)
	}

	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}
