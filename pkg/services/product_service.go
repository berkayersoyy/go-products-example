package services

import (
	"github.com/berkay.ersoyy/go-products-example/pkg/models"
	"github.com/berkay.ersoyy/go-products-example/pkg/repositories"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func ProvideProductService(p repositories.ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) GetAllProducts() []models.Product {
	return p.ProductRepository.GetAllProducts()
}

func (p *ProductService) GetProductByID(id uint) models.Product {
	return p.ProductRepository.GetProductByID(id)
}

func (p *ProductService) AddProduct(product models.Product) models.Product {
	p.ProductRepository.AddProduct(product)

	return product
}

func (p *ProductService) DeleteProduct(product models.Product) {
	p.ProductRepository.DeleteProduct(product)
}
