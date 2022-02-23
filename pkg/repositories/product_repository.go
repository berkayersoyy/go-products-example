package repositories

import (
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	GetAllProducts() []models.Product
	GetProductByID(id uint) models.Product
	AddProduct(product models.Product) models.Product
	DeleteProduct(product models.Product)
}

type productRepository struct {
	DB *gorm.DB
}

func ProvideProductRepository(DB *gorm.DB) ProductRepository {
	return &productRepository{DB: DB}
}

func (p *productRepository) GetAllProducts() []models.Product {
	var products []models.Product
	p.DB.Find(&products)

	return products
}

func (p *productRepository) GetProductByID(id uint) models.Product {
	var product models.Product
	p.DB.First(&product, id)

	return product
}

func (p *productRepository) AddProduct(product models.Product) models.Product {
	p.DB.Save(&product)

	return product
}

func (p *productRepository) DeleteProduct(product models.Product) {
	p.DB.Delete(&product)
}
