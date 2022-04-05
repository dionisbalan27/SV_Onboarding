package product_usecase

import (
	"backend-api/helpers"
	"backend-api/models/product/dto"
	"backend-api/models/product/entity"
	"backend-api/repository/product_repository"
	"errors"

	"gorm.io/gorm"
)

type ProductUsecase interface {
	GetAllProducts() dto.Response
	GetProductById(string) dto.Response
	CreateNewProduct(dto.Product, string) dto.Response
	UpdateProductData(dto.Product, string) dto.Response
	UpdatePublishProduct(dto.Product, string, string) dto.Response
	UpdateCheckProduct(dto.Product, string, string) dto.Response
	DeleteProductById(string) dto.Response
}

type productUsecase struct {
	productRepo product_repository.ProductRepository
}

type ProductUcaseTest struct {
	productRepo *product_repository.ProductRepositoryMock
}

func GetProductUsecase(productRepository product_repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepository,
	}
}

func (product *productUsecase) GetAllProducts() dto.Response {
	productlist, err := product.productRepo.GetAllProducts()

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err.Error(), 500)
	}

	return helpers.ResponseSuccess("ok", nil, response, 200)
}

func (product *productUsecase) CreateNewProduct(newProduct dto.Product, id_user string) dto.Response {

	productInsert := entity.Product{
		ID:          newProduct.ID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
	}

	productData, err := product.productRepo.CreateNewProduct(productInsert, id_user)

	if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": productData.ID}, 201)
}

func (product *productUsecase) UpdateProductData(productUpdate dto.Product, id string) dto.Response {

	productInsert := convertToProductEntity(productUpdate)

	_, err := product.productRepo.UpdateProductData(productInsert, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	productUpdate.ID = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (product *productUsecase) UpdateCheckProduct(productUpdate dto.Product, id string, id_user string) dto.Response {

	productInsert := convertToProductEntity(productUpdate)

	_, err := product.productRepo.UpdateCheckProduct(productInsert, id, id_user)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	productUpdate.ID = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (product *productUsecase) UpdatePublishProduct(productUpdate dto.Product, id string, id_user string) dto.Response {

	productInsert := convertToProductEntity(productUpdate)

	_, err := product.productRepo.UpdatePublishProduct(productInsert, id, id_user)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	productUpdate.ID = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (product *productUsecase) DeleteProductById(id string) dto.Response {

	err := product.productRepo.DeleteProductById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	return helpers.ResponseSuccess("ok", nil, nil, 200)
}

func (product *productUsecase) GetProductById(id string) dto.Response {
	userData, err := product.productRepo.GetProductById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}

	return helpers.ResponseSuccess("ok", nil, userData, 200)
}

func convertToProductEntity(productUpdate dto.Product) entity.Product {
	return entity.Product{
		Name:        productUpdate.Name,
		Description: productUpdate.Description,
	}
}
