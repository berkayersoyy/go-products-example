package repositories

import (
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}
func (u *UserRepository) GetAllUsers() []models.User {
	var users []models.User
	u.DB.Find(&users)

	return users
}

func (u *UserRepository) GetUserByID(id uint) models.User {
	var user models.User
	u.DB.First(&user, id)

	return user
}
func (u *UserRepository) GetUserByUsername(username string) models.User {
	var user models.User
	u.DB.Where("username = ?", username).First(&user)

	return user
}

func (u *UserRepository) AddUser(user models.User) models.User {
	u.DB.Save(&user)

	return user
}

func (u *UserRepository) DeleteUser(user models.User) {
	u.DB.Delete(&user)
}
