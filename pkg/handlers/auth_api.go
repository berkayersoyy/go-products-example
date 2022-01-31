package handlers

import (
	"net/http"

	"github.com/berkay.ersoyy/go-products-example/pkg/models"
	"github.com/berkay.ersoyy/go-products-example/pkg/services"
	"github.com/gin-gonic/gin"
)

type AuthAPI struct {
	AuthService services.AuthService
	UserService services.UserService
}

func (a *AuthAPI) Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	user := a.UserService.GetUserByUsername(u.Username)
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := a.AuthService.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
