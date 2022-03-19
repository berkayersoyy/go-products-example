// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	models "github.com/berkayersoyy/go-products-example/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: product
func (_m *ProductRepository) AddProduct(product models.Product) models.Product {
	ret := _m.Called(product)

	var r0 models.Product
	if rf, ok := ret.Get(0).(func(models.Product) models.Product); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	return r0
}

// DeleteProduct provides a mock function with given fields: product
func (_m *ProductRepository) DeleteProduct(product models.Product) {
	_m.Called(product)
}

// GetAllProducts provides a mock function with given fields:
func (_m *ProductRepository) GetAllProducts() []models.Product {
	ret := _m.Called()

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func() []models.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	return r0
}

// GetProductByID provides a mock function with given fields: id
func (_m *ProductRepository) GetProductByID(id uint) models.Product {
	ret := _m.Called(id)

	var r0 models.Product
	if rf, ok := ret.Get(0).(func(uint) models.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	return r0
}
