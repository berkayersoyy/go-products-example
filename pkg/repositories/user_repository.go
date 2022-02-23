package repositories

import (
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	DB *gorm.DB
}
type UserRepository interface {
	GetAllUsers() []models.User
	GetUserByID(id uint) models.User
	GetUserByUsername(username string) models.User
	AddUser(user models.User) models.User
	DeleteUser(user models.User)
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{DB: DB}
}
func (u *userRepository) GetAllUsers() []models.User {
	var users []models.User
	u.DB.Find(&users)

	return users
}

func (u *userRepository) GetUserByID(id uint) models.User {
	var user models.User
	u.DB.First(&user, id)

	return user
}
func (u *userRepository) GetUserByUsername(username string) models.User {
	var user models.User
	u.DB.Where("username = ?", username).First(&user)

	return user
}

func (u *userRepository) AddUser(user models.User) models.User {
	u.DB.Save(&user)

	return user
}

func (u *userRepository) DeleteUser(user models.User) {
	u.DB.Delete(&user)
}
