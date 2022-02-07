package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/berkayersoyy/go-products-example/pkg/dto"
	"github.com/berkayersoyy/go-products-example/pkg/mappers"
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/berkayersoyy/go-products-example/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAPI struct {
	UserService services.UserService
}

func ProvideUserAPI(u services.UserService) UserAPI {
	return UserAPI{UserService: u}
}

// @BasePath /api/v1

// GetAllUsers
// @Summary Fetch all users
// @Schemes
// @Description Fetch all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/ [get]
func (u *UserAPI) GetAllUsers(c *gin.Context) {
	users := u.UserService.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// @BasePath /api/v1

// GetUserByID
// @Summary Fetch user by id
// @Schemes
// @Description Fetch user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/{id} [get]
func (u *UserAPI) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.GetUserByID(uint(id))
	if user == (models.User{}) {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": mappers.ToUserDTO(user)})
}

// @BasePath /api/v1

// AddUser
// @Summary Add user
// @Schemes
// @Description Add user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/ [post]
func (u *UserAPI) AddUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	createdUser := u.UserService.AddUser(user)
	c.JSON(http.StatusOK, gin.H{"user": mappers.ToUserDTO(createdUser)})
}

// @BasePath /api/v1

// Update User
// @Summary Update user
// @Schemes
// @Description Update user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.UserDTO true "User ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/ [put]
func (u *UserAPI) UpdateUser(c *gin.Context) {
	var userDTO dto.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.GetUserByID(uint(id))
	if user == (models.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	user.Username = userDTO.Username
	user.Password = userDTO.Password
	u.UserService.AddUser(user)

	c.Status(http.StatusOK)
}

// @BasePath /api/v1

// DeleteUser
// @Summary Delete user
// @Schemes
// @Description Delete user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/{id} [delete]
func (u *UserAPI) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.GetUserByID(uint(id))
	if user == (models.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	u.UserService.DeleteUser(user)
	c.Status(http.StatusOK)
}
func (u *UserAPI) GetUserByUsername(c *gin.Context) {

	un := c.Param("username")
	user := u.UserService.GetUserByUsername(un)
	c.JSON(http.StatusOK, gin.H{"user": mappers.ToUserDTO(user)})

}
