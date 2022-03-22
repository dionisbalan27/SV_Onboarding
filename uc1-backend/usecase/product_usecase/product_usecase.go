package product_usecase

import (
	helpers "backend-api/helpers/helpers_product"
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

func GetProductUsecase(productRepository product_repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepository,
	}
}

func (product *productUsecase) GetAllProducts() dto.Response {
	productlist, err := product.productRepo.GetAllProducts()

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Error Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}
	return helpers.ResponseSuccess("ok", nil, productlist)
}

func (product *productUsecase) CreateNewProduct(newProduct dto.Product, id_user string) dto.Response {

	productInsert := entity.Product{
		ID:          newProduct.ID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
	}

	productData, err := product.productRepo.CreateNewProduct(productInsert, id_user)

	if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": productData.ID})
}

func (product *productUsecase) UpdateProductData(productUpdate dto.Product, id string) dto.Response {

	productInsert := convertToProductEntity(productUpdate)

	_, err := product.productRepo.UpdateProductData(productInsert, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Error Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	productUpdate.ID = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id})
}

func (product *productUsecase) UpdateCheckProduct(productUpdate dto.Product, id string, id_user string) dto.Response {

	productInsert := convertToProductEntity(productUpdate)

	_, err := product.productRepo.UpdateCheckProduct(productInsert, id, id_user)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Error Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	productUpdate.ID = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id})
}

func (product *productUsecase) UpdatePublishProduct(productUpdate dto.Product, id string, id_user string) dto.Response {

	productInsert := convertToProductEntity(productUpdate)

	_, err := product.productRepo.UpdatePublishProduct(productInsert, id, id_user)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Error Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	productUpdate.ID = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id})
}

func (product *productUsecase) DeleteProductById(id string) dto.Response {

	err := product.productRepo.DeleteProductById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Error Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("User deleted successfully", 200, nil)
}

func (product *productUsecase) GetProductById(id string) dto.Response {
	userData, err := product.productRepo.GetProductById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Error Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}

	return helpers.ResponseSuccess("ok", nil, userData)
}

func convertToProductEntity(productUpdate dto.Product) entity.Product {
	return entity.Product{
		Name:        productUpdate.Name,
		Description: productUpdate.Description,
	}
}
