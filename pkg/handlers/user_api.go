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
// @Summary Fetch all users from database
// @Schemes
// @Description Fetch all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /v1/users/ [get]
func (u *UserAPI) GetAllUsers(c *gin.Context) {
	users := u.UserService.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func (u *UserAPI) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.GetUserByID(uint(id))
	c.JSON(http.StatusOK, gin.H{"user": mappers.ToUserDTO(user)})
}
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
