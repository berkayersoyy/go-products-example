// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	models "github.com/berkayersoyy/go-products-example/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: user
func (_m *UserRepository) AddUser(user models.User) models.User {
	ret := _m.Called(user)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(models.User) models.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: user
func (_m *UserRepository) DeleteUser(user models.User) {
	_m.Called(user)
}

// GetAllUsers provides a mock function with given fields:
func (_m *UserRepository) GetAllUsers() []models.User {
	ret := _m.Called()

	var r0 []models.User
	if rf, ok := ret.Get(0).(func() []models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	return r0
}

// GetUserByID provides a mock function with given fields: id
func (_m *UserRepository) GetUserByID(id uint) models.User {
	ret := _m.Called(id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(uint) models.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	return r0
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *UserRepository) GetUserByUsername(username string) models.User {
	ret := _m.Called(username)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	return r0
}
