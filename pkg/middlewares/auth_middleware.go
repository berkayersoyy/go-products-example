package middlewares

import (
	"fmt"
	"net/http"

	"github.com/berkayersoyy/go-products-example/pkg/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWTMiddleware(a services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// const BEARER_SCHEMA = "Bearer"
		// authHeader := c.GetHeader("Authorization")
		// tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		token, err := a.ValidateToken(c.Request)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
