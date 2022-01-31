package validators

import (
	"net/http"

	"github.com/berkay.ersoyy/go-products-example/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.ShouldBindJSON(&user)
		if err == nil {
			validate := validator.New()
			err := validate.Struct(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
