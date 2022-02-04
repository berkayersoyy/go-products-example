package services

import (
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/berkayersoyy/go-products-example/pkg/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func ProvideUserService(u repositories.UserRepository) UserService {
	return UserService{UserRepository: u}
}
func (u *UserService) GetAllUsers() []models.User {
	return u.UserRepository.GetAllUsers()
}
func (u *UserService) GetUserByID(id uint) models.User {
	return u.UserRepository.GetUserByID(id)
}
func (u *UserService) AddUser(user models.User) models.User {
	return u.UserRepository.AddUser(user)
}
func (u *UserService) GetUserByUsername(username string) models.User {
	return u.UserRepository.GetUserByUsername(username)
}
func (u *UserService) DeleteUser(user models.User) {
	u.UserRepository.DeleteUser(user)
}
