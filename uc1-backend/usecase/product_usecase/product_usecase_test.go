package product_usecase

import (
	"backend-api/models/product/dto"
	"backend-api/models/product/entity"
	"backend-api/repository/product_repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = product_repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productUc = ProductUcaseTest{productRepo: &productRepository}

func TestGetAllProductsSuccess(t *testing.T) {
	expected := []entity.Product{}
	productRepository.Mock.On("GetAllProducts").Return(expected, nil)
	result, err := productUc.productRepo.GetAllProducts()

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestGetProductByIdNotFound(t *testing.T) {
	productRepository.Mock.On("GetProductById", "1").Return(nil, errors.New("Data not found"))

	result, err := productUc.productRepo.GetProductById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestGetProductByIdSuccess(t *testing.T) {
	expected := &entity.Product{}

	productRepository.Mock.On("GetProductById", "2").Return(expected, nil)

	result, err := productUc.productRepo.GetProductById("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestUpdateProductNotFound(t *testing.T) {
	productRepository.Mock.On("GetProductById", "1").Return(nil, errors.New("Data not found"))

	result, err := productUc.productRepo.GetProductById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestUpdateProductSuccess(t *testing.T) {
	expected := &entity.Product{}

	productRepository.Mock.On("UpdateProductData", "2").Return(expected, nil)

	result, err := productUc.productRepo.UpdateProductData("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestCreateNewProductSuccess(t *testing.T) {
	request := dto.Product{
		Name:        "New Product",
		Description: "Test New Product",
	}

	expected := &entity.Product{}

	productRepository.Mock.On("CreateNewProduct").Return(expected, nil)

	result, err := productUc.productRepo.CreateNewProduct(request)

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestCreateNewProductSFailed(t *testing.T) {
	request := dto.Product{
		Name:        "New Product",
		Description: "Test New Product",
	}

	expected := &entity.Product{}

	productRepository.Mock.On("CreateNewProduct").Return(nil, errors.New("Failed create product"))

	result, err := productUc.productRepo.CreateNewProduct(request)

	assert.NotEqual(t, expected, result)
	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestDeleteProductFailed(t *testing.T) {
	productRepository.Mock.On("DeleteProductById", "1").Return(nil, errors.New("Data not found"))

	result, err := productUc.productRepo.DeleteProductById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestDeleteProductSuccess(t *testing.T) {
	expected := &entity.Product{}

	productRepository.Mock.On("DeleteProductById", "2").Return(expected, nil)

	result, err := productUc.productRepo.DeleteProductById("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}
