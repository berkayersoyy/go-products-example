package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/berkayersoyy/go-products-example/pkg/database"
	"github.com/berkayersoyy/go-products-example/pkg/dto"
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/berkayersoyy/go-products-example/pkg/repositories"
	"github.com/berkayersoyy/go-products-example/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *UserSuite) SetupSuite() {
	suite.db = database.GetMysqlClient("../../").SingletonMysql
}
func (suite *UserSuite) TearDownTest() {

}

func (suite *UserSuite) TearDownSuite() {
	defer suite.db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
func (suite *UserSuite) TestGetAllUsers() {
	req, w := setGetAllUsersRouter(suite.db)
	a := suite.Assert()

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.User{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.Model = gorm.Model{}
	expected := models.User{}
	a.Equal(expected, actual)

}
func setGetAllUsersRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	api := ProvideUserAPI(services.ProvideUserService(repositories.ProvideUserRepository(db)))
	r.GET("/v1/users", api.GetAllUsers)
	req, err := http.NewRequest(http.MethodGet, "/v1/users", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w

}
func (suite *UserSuite) TestGetUserByID() {
	req, w := setGetUserByIDRouter(suite.db, "/v1/users/1")

	a := suite.Assert()

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.User{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.Model = gorm.Model{}
	expected := models.User{}
	a.Equal(expected, actual)

}
func setGetUserByIDRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	api := ProvideUserAPI(services.ProvideUserService(repositories.ProvideUserRepository(db)))
	r.GET("/v1/users/:id", api.GetUserByID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w

}

func (suite *UserSuite) TestAddUser() {
	a := suite.Assert()
	user := models.User{
		Username: "test-username",
		Password: "test-password",
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setAddUserRouter(suite.db, bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.User{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.Model = gorm.Model{}
	expected := models.User{}
	a.Equal(expected, actual)

}
func setAddUserRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	api := ProvideUserAPI(services.ProvideUserService(repositories.ProvideUserRepository(db)))
	r.POST("/v1/users", api.AddUser)
	req, err := http.NewRequest(http.MethodPost, "/v1/users", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}
func (suite *UserSuite) TestUpdateUser() {
	a := suite.Assert()
	user := dto.UserDTO{
		ID:       1,
		Username: "test-username-changed",
		Password: "test-password-changed",
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setUpdateUserRouter(suite.db, bytes.NewBuffer(reqBody), "/v1/users/1")

	a.Equal(http.MethodPut, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.User{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.Model = gorm.Model{}
	expected := models.User{}
	a.Equal(expected, actual)

}
func setUpdateUserRouter(db *gorm.DB, body *bytes.Buffer, url string) (*http.Request, *httptest.ResponseRecorder, error) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	api := ProvideUserAPI(services.ProvideUserService(repositories.ProvideUserRepository(db)))
	r.PUT("/v1/users/:id", api.UpdateUser)
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func (suite *UserSuite) TestDeleteUser() {
	a := suite.Assert()

	req, w, err := setDeleteUserRouter(suite.db, "/v1/users/1")

	if err != nil {
		a.Error(err)
	}
	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

}
func setDeleteUserRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder, error) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	api := ProvideUserAPI(services.ProvideUserService(repositories.ProvideUserRepository(db)))
	r.DELETE("/v1/users/:id", api.DeleteUser)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}
