package mappers

import (
	"github.com/berkayersoyy/go-products-example/pkg/dto"
	"github.com/berkayersoyy/go-products-example/pkg/models"
)

func ToProduct(productDTO dto.ProductDTO) models.Product {

	return models.Product{Name: productDTO.Name, Price: productDTO.Price, Description: productDTO.Description}
}

func ToProductDTO(product models.Product) dto.ProductDTO {
	return dto.ProductDTO{ID: product.ID, Price: product.Price, Name: product.Name, Description: product.Description}
}

func ToProductDTOs(products []models.Product) []dto.ProductDTO {
	productdtos := make([]dto.ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}
