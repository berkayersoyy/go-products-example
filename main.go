package main

import (
	_ "github.com/berkayersoyy/go-products-example/docs"
	"github.com/berkayersoyy/go-products-example/pkg/database"
	"github.com/berkayersoyy/go-products-example/pkg/handlers"
	"github.com/berkayersoyy/go-products-example/pkg/middlewares"
	"github.com/berkayersoyy/go-products-example/pkg/repositories"
	"github.com/berkayersoyy/go-products-example/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func setup(db *gorm.DB) *gin.Engine {

	productRepository := repositories.ProvideProductRepository(db)
	productService := services.ProvideProductService(productRepository)
	productApi := handlers.ProvideProductAPI(productService)

	userRepository := repositories.ProvideUserRepository(db)
	userService := services.ProvideUserService(userRepository)
	userApi := handlers.ProvideUserAPI(userService)

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	db := database.GetMysqlClient()
	defer db.SingletonMysql.Close()
	r := setup(db.SingletonMysql)
	err := r.Run()
	if err != nil {
		panic(err)
	}

}
