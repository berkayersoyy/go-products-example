package repositories

import (
	"github.com/berkay.ersoyy/go-products-example/pkg/models"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepository(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) GetAllProducts() []models.Product {
	var products []models.Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) GetProductByID(id uint) models.Product {
	var product models.Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) AddProduct(product models.Product) models.Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) DeleteProduct(product models.Product) {
	p.DB.Delete(&product)
}
