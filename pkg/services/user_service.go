package services

import (
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/berkayersoyy/go-products-example/pkg/repositories"
)

type userService struct {
	UserRepository repositories.UserRepository
}
type UserService interface {
	GetAllUsers() []models.User
	GetUserByID(id uint) models.User
	AddUser(user models.User) models.User
	GetUserByUsername(username string) models.User
	DeleteUser(models.User)
}

func ProvideUserService(u repositories.UserRepository) UserService {
	return &userService{UserRepository: u}
}
func (u *userService) GetAllUsers() []models.User {
	return u.UserRepository.GetAllUsers()
}
func (u *userService) GetUserByID(id uint) models.User {
	return u.UserRepository.GetUserByID(id)
}
func (u *userService) AddUser(user models.User) models.User {
	return u.UserRepository.AddUser(user)
}
func (u *userService) GetUserByUsername(username string) models.User {
	return u.UserRepository.GetUserByUsername(username)
}
func (u *userService) DeleteUser(user models.User) {
	u.UserRepository.DeleteUser(user)
}
