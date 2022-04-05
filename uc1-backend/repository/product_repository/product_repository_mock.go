package product_repository

import (
	"backend-api/models/product/dto"
	"backend-api/models/product/entity"
	"errors"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) GetAllProducts() ([]entity.Product, error) {
	ret := repository.Mock.Called()

	var r0 []entity.Product
	if rf, ok := ret.Get(0).(func() []entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (repository *ProductRepositoryMock) GetProductById(id string) (*entity.Product, error) {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil, errors.New("Data not found")
	} else {
		product := args.Get(0).(*entity.Product)
		return product, nil
	}
}

func (repository *ProductRepositoryMock) UpdateProductData(id string) (*entity.Product, error) {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil, errors.New("Data not found")
	} else {
		product := args.Get(0).(*entity.Product)
		return product, nil
	}
}

func (repository *ProductRepositoryMock) DeleteProductById(id string) (*entity.Product, error) {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil, errors.New("Data not found")
	} else {
		product := args.Get(0).(*entity.Product)
		return product, nil
	}
}

func (repository *ProductRepositoryMock) CreateNewProduct(newProduct dto.Product) (*entity.Product, error) {
	args := repository.Mock.Called()

	if args.Get(0) == nil {
		return nil, errors.New("Failed create product")
	} else {

		product := args.Get(0).(*entity.Product)

		return product, nil
	}
}
