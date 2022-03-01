package repositories

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	productRepository ProductRepository
	userRepository    UserRepository
	product           *models.Product
	user              *models.User
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("sqlmock", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.mock.MatchExpectationsInOrder(false)
	s.productRepository = ProvideProductRepository(s.DB)
	s.userRepository = ProvideUserRepository(s.DB)
}
func (s *Suite) TestRepositoryGetProductById() {
	product := models.Product{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, Name: "test-product", Price: 10, Description: "test-description"}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products`  WHERE `products`.`deleted_at` IS NULL AND ((`products`.`id` = 1)) ORDER BY `products`.`id` ASC LIMIT 1")).
		WithArgs(product.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "price", "description"}).
			AddRow(product.ID, product.CreatedAt, product.UpdatedAt, product.DeletedAt, product.Name, product.Price, product.Description))
	res := s.productRepository.GetProductByID(product.ID)

	require.Equal(s.T(), product, res)
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}
}

func (s *Suite) TestRepositoryGetAllProducts() {
	product := models.Product{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, Name: "test-product", Price: 10, Description: "test-description"}
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products`  WHERE `products`.`deleted_at` IS NULL ORDER BY `products`.`id` ASC LIMIT 1")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "price", "description"}).
			AddRow(product.ID, product.CreatedAt, product.UpdatedAt, product.DeletedAt, product.Name, product.Price, product.Description))
	res := s.productRepository.GetAllProducts()
	require.Equal(s.T(), product, res[0])
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}
}
func (s *Suite) TestRepositoryAddProduct() {
	product := models.Product{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, Name: "test-product", Price: 10, Description: "test-description"}
	prep := s.mock.ExpectPrepare("INSERT INTO products (name, price, description) VALUES (?, ?, ?)")
	prep.ExpectExec().
		WithArgs(product.Name, product.Price, product.Description).
		WillReturnResult(sqlmock.NewResult(0, 1))
	res := s.productRepository.AddProduct(product)
	require.Equal(s.T(), product, res)
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}
}
func (s *Suite) TestRepositoryDeleteProduct() {
	product := models.Product{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, Name: "test-product", Price: 10, Description: "test-description"}
	prep := s.mock.ExpectPrepare("DELETE from products WHERE id = ?")
	prep.ExpectExec().
		WithArgs(product.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.productRepository.DeleteProduct(product)
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}
}
func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
