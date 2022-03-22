package product_repository

import (
	"backend-api/models/product/entity"

	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.ProductList, error)
	GetProductById(string) (*entity.ProductDetailID, error)
	CreateNewProduct(entity.Product, string) (*entity.Product, error)
	UpdateProductData(entity.Product, string) (*entity.Product, error)
	UpdateCheckProduct(entity.Product, string, string) (*entity.Product, error)
	UpdatePublishProduct(entity.Product, string, string) (*entity.Product, error)
	DeleteProductById(string) error
}

type productRepository struct {
	mysqlConnection *gorm.DB
}

func GetProductRepository(mysqlConn *gorm.DB) ProductRepository {
	return &productRepository{
		mysqlConnection: mysqlConn,
	}
}

func (repo *productRepository) GetAllProducts() ([]entity.ProductList, error) {

	products := []entity.ProductList{}
	err := repo.mysqlConnection.Model(&entity.Product{}).Select("products.id, products.name, products.description, products.status").Scan(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return products, nil
}

func (repo *productRepository) CreateNewProduct(product entity.Product, id_user string) (*entity.Product, error) {
	product.ID = uuid.New().String()
	product.Status = "inactive"
	product.MakerID = id_user

	if err := repo.mysqlConnection.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepository) GetProductById(id string) (*entity.ProductDetailID, error) {
	product := entity.ProductDetail{}
	users := []entity.Action{}

	// err := repo.mysqlConnection.Model(&entity.Product{}).Where("id = ?", id).Find(&product).Error

	// dapatkan product berdasarkan Id produk
	err := repo.mysqlConnection.Model(&entity.Product{}).Where("products.id = ?", id).Select("products.id, products.name, products.description, products.status, products.maker_id, products.signer_id,  products.checker_id").Scan(&product).Error
	if err != nil {
		return nil, err
	}

	// ambil makerId, checkerId dan signerId
	// lakukan query ke tabel users untuk mendapatkan data maker, checker dan signer WHERE IN
	err = repo.mysqlConnection.Model(&entity.User{}).Where("users.id IN ?", []string{product.MakerID, product.CheckerID, product.SignerID}).Select("users.id, users.name").Find(&users).Error
	if err != nil {
		return nil, err
	}

	Action := entity.Action{
		ID:   "",
		Name: "",
	}

	productDetail := entity.ProductDetailID{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Status:      product.Status,
		Maker:       Action,
		Checker:     Action,
		Signer:      Action,
	}

	// setelah itu ambil data sesuai dengan id lalu masukkan ke data produk sesuai dengan makerId, checkerId dan signerId
	for _, usr := range users {
		if usr.ID == product.MakerID {
			productDetail.Maker.ID = usr.ID
			productDetail.Maker.Name = usr.Name

		} else if usr.ID == product.CheckerID {
			productDetail.Checker.ID = usr.ID
			productDetail.Checker.Name = usr.Name
		} else if usr.ID == product.SignerID {
			productDetail.Signer.ID = usr.ID
			productDetail.Signer.Name = usr.Name
		}
	}

	return &productDetail, nil
}

func (repo *productRepository) UpdateProductData(product entity.Product, id string) (*entity.Product, error) {

	if err := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description}).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepository) UpdateCheckProduct(product entity.Product, id string, id_user string) (*entity.Product, error) {

	if err := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description, "checker_id": id_user, "status": "registered"}).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepository) UpdatePublishProduct(product entity.Product, id string, id_user string) (*entity.Product, error) {

	if err := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description, "signer_id": id_user, "status": "active"}).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepository) DeleteProductById(id string) error {
	sql := "DELETE FROM products"

	sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)

	if err := repo.mysqlConnection.Raw(sql).Scan(entity.Product{}).Error; err != nil {

		return err
	}
	// if err := repo.mysqlConnection.Delete(&entity.User{}, id).Error; err != nil  {
	// 	return err
	// }
	return nil
}
