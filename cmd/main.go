package main

import (
	"github.com/berkay.ersoyy/go-products-example/pkg/database"
	"github.com/berkay.ersoyy/go-products-example/pkg/handlers"
	"github.com/berkay.ersoyy/go-products-example/pkg/middlewares"
	"github.com/berkay.ersoyy/go-products-example/pkg/repositories"
	"github.com/berkay.ersoyy/go-products-example/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setup(db *gorm.DB) *gin.Engine {

	productRepository := repositories.ProductRepository{DB: db}
	productService := services.ProductService{ProductRepository: productRepository}
	productApi := handlers.ProductAPI{ProductService: productService}

	userRepository := repositories.UserRepository{DB: db}
	userService := services.UserService{UserRepository: userRepository}
	userApi := handlers.UserAPI{UserService: userService}

	authService := services.AuthService{}
	authApi := handlers.AuthAPI{AuthService: authService, UserService: userService}

	//TODO Dependency injection with wire.go but
	// productApi := InitProductAPI(db)

	router := gin.Default()

	//TODO Middleware for validation
	//router.Use(validators.ProductValidator())

	//TODO Error handler can be add as a middleware
	//TODO Swagger

	//products
	products := router.Group("/v1")

	products.Use(middlewares.AuthorizeJWTMiddleware(authService))

	products.GET("/products", productApi.GetAllProducts)
	products.POST("/products", productApi.AddProduct)
	products.GET("/products/:id", productApi.GetProductByID)
	products.DELETE("/products/:id", productApi.DeleteProduct)
	products.PUT("/products/:id", productApi.UpdateProduct)

	//users
	users := router.Group("/v1")
	users.GET("/users", userApi.GetAllUsers)
	users.POST("/users", userApi.AddUser)
	users.GET("/users/:id", userApi.GetUserByID)
	users.DELETE("/users/:id", userApi.DeleteUser)
	users.PUT("/users/:id", userApi.UpdateUser)

	//auth
	auth := router.Group("/v1")
	auth.POST("/login", authApi.Login)

	return router
}

func main() {
	db := database.InitDb()
	defer db.Close()
	r := setup(db)
	err := r.Run()
	if err != nil {
		panic(err)
	}

}
