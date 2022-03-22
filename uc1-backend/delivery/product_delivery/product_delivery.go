package product_delivery

import (
	helpers "backend-api/helpers/helpers_product"
	"backend-api/models/product/dto"
	"fmt"
	"net/http"

	"backend-api/usecase/product_usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductDelivery interface {
	GetAllProducts(*gin.Context)
	CreateNewProduct(*gin.Context)
	UpdateProductData(*gin.Context)
	DeleteProductById(c *gin.Context)
	GetProductById(c *gin.Context)
	UpdateCheckProduct(c *gin.Context)
	UpdatePublishProduct(c *gin.Context)
}

type productDelivery struct {
	usecase product_usecase.ProductUsecase
}

func GetProductDelivery(productUsecase product_usecase.ProductUsecase) ProductDelivery {
	return &productDelivery{
		usecase: productUsecase,
	}
}

func (res *productDelivery) GetAllProducts(c *gin.Context) {
	response := res.usecase.GetAllProducts()
	// fmt.Printf("%+v", response)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *productDelivery) GetProductById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.GetProductById(id)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *productDelivery) CreateNewProduct(c *gin.Context) {
	request := dto.Product{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Invalid Input", 400)
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}
	}
	id_user, _ := c.Get("user_id")
	str_id_user := fmt.Sprintf("%v", id_user)
	response := res.usecase.CreateNewProduct(request, str_id_user)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (res *productDelivery) UpdateProductData(c *gin.Context) {
	id := c.Param("id")
	request := dto.Product{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Invalid Input", 400)
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}

	}
	response := res.usecase.UpdateProductData(request, id)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *productDelivery) UpdateCheckProduct(c *gin.Context) {
	id := c.Param("id")
	request := dto.Product{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Invalid Input", 400)
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}

	}
	id_user, _ := c.Get("user_id")
	str_id_user := fmt.Sprintf("%v", id_user)
	response := res.usecase.UpdateCheckProduct(request, id, str_id_user)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *productDelivery) UpdatePublishProduct(c *gin.Context) {
	id := c.Param("id")
	request := dto.Product{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Invalid Input", 400)
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}

	}
	id_user, _ := c.Get("user_id")
	str_id_user := fmt.Sprintf("%v", id_user)
	response := res.usecase.UpdatePublishProduct(request, id, str_id_user)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *productDelivery) DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteProductById(id)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.JSON(http.StatusOK, response)
}
