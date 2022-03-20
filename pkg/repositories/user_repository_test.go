package repositories

import (
	mocks "github.com/berkayersoyy/go-products-example/pkg/mocks/repositories"
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_GetAllUsersShouldReturnNotEmptyUserArray(t *testing.T) {
	users := []models.User{{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("GetAllUsers").Return(users)

	resp := mockRepo.GetAllUsers()

	assert.Equal(t, users, resp)
	assert.NotEmpty(t, users)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetAllUsers", 1)
}
func TestUserRepository_GetAllUsersShouldReturnEmptyUserArray(t *testing.T) {
	users := []models.User{}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("GetAllUsers").Return(users)

	resp := mockRepo.GetAllUsers()

	assert.Equal(t, users, resp)
	assert.Empty(t, users)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetAllUsers", 1)
}
func TestUserRepository_GetUserByIDShouldReturnValidUser(t *testing.T) {
	user := models.User{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("GetUserByID", user.ID).Return(user)

	resp := mockRepo.GetUserByID(uint(1))

	assert.Equal(t, user, resp)
	assert.NotEmpty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetUserByID", 1)
}
func TestUserRepository_GetUserByIDShouldReturnEmptyUser(t *testing.T) {
	user := models.User{}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("GetUserByID", uint(0)).Return(user)

	resp := mockRepo.GetUserByID(uint(0))

	assert.Equal(t, user, resp)
	assert.Empty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetUserByID", 1)
}
func TestUserRepository_GetUserByUsernameShouldReturnValidUser(t *testing.T) {
	user := models.User{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("GetUserByUsername", user.Username).Return(user)

	resp := mockRepo.GetUserByUsername("test-username")

	assert.Equal(t, user, resp)
	assert.NotEmpty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetUserByUsername", 1)
}
func TestUserRepository_GetUserByUsernameShouldReturnEmptyUser(t *testing.T) {
	user := models.User{}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("GetUserByUsername", "test").Return(user)

	resp := mockRepo.GetUserByUsername("test")

	assert.Equal(t, user, resp)
	assert.Empty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetUserByUsername", 1)
}
func TestUserRepository_AddUserShouldReturnValidUser(t *testing.T) {
	user := models.User{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("AddUser", user).Return(user)

	resp := mockRepo.AddUser(user)

	assert.Equal(t, user, resp)
	assert.NotEmpty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "AddUser", 1)
}
func TestUserRepository_AddUserShouldReturnEmptyUser(t *testing.T) {
	user := models.User{}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("AddUser", user).Return(user)

	resp := mockRepo.AddUser(user)

	assert.Equal(t, user, resp)
	assert.Empty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "AddUser", 1)
}
func TestUserRepository_DeleteUser(t *testing.T) {
	user := models.User{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}
	mockRepo := mocks.UserRepository{}
	mockRepo.On("DeleteUser", user)

	mockRepo.DeleteUser(user)

	mockRepo.AssertNumberOfCalls(t, "AddUser", 1)
}
