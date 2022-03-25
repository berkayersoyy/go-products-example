package handlers

import (
	"bytes"
	"encoding/json"
	mocks "github.com/berkayersoyy/go-products-example/pkg/mocks/handlers"
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserAPI_GetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		users := []models.User{{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}}
		mockApi := mocks.UserAPI{}
		mockApi.On("GetAllUsers", mock.Anything).Return(users)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.GET("/v1/users", mockApi.GetAllUsers)
		req, err := http.NewRequest(http.MethodGet, "/v1/users", nil)
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)
		r.ServeHTTP(w, req)

		var respBody []models.User
		err = json.NewDecoder(w.Body).Decode(&respBody)
		assert.NoError(t, err)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, users, respBody)
		mockApi.AssertNumberOfCalls(t, "GetAllUsers", 1)
	})

}
func TestUserAPI_GetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		users := []models.User{{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}}
		mockApi := mocks.UserAPI{}
		mockApi.On("GetUserByID", mock.Anything).Return(users)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.Use(func(context *gin.Context) {
			context.Set("id", uint(1))
		})
		r.GET("/v1/users/:id", mockApi.GetUserByID)
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)
		r.ServeHTTP(w, req)

		var respBody []models.User
		err = json.NewDecoder(w.Body).Decode(&respBody)
		assert.NoError(t, err)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, users, respBody)
		mockApi.AssertNumberOfCalls(t, "GetUserByID", 1)
	})

}
func TestUserAPI_AddUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		user := models.User{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}
		mockApi := mocks.UserAPI{}
		mockApi.On("AddUser", mock.Anything).Return(user)

		w := httptest.NewRecorder()
		r := gin.Default()

		reqBody, err := json.Marshal(user)
		assert.NoError(t, err)
		r.POST("/v1/users", mockApi.AddUser)
		req, err := http.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)
		r.ServeHTTP(w, req)

		var respBody models.User
		err = json.NewDecoder(w.Body).Decode(&respBody)
		assert.NoError(t, err)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, user, respBody)
		mockApi.AssertNumberOfCalls(t, "AddUser", 1)
	})
}
func TestUserAPI_UpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		user := models.User{Username: "test-username", Password: "test-pass", Model: gorm.Model{ID: 1}}
		mockApi := mocks.UserAPI{}
		mockApi.On("UpdateUser", mock.Anything).Return(user)

		w := httptest.NewRecorder()
		r := gin.Default()

		reqBody, err := json.Marshal(user)
		assert.NoError(t, err)
		r.PUT("/v1/users", mockApi.UpdateUser)
		req, err := http.NewRequest(http.MethodPut, "/v1/users", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)
		r.ServeHTTP(w, req)

		var respBody models.User
		err = json.NewDecoder(w.Body).Decode(&respBody)
		assert.NoError(t, err)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, user, respBody)
		mockApi.AssertNumberOfCalls(t, "UpdateUser", 1)
	})
}
func TestUserAPI_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		mockApi := mocks.UserAPI{}
		mockApi.On("DeleteUser", mock.Anything)

		w := httptest.NewRecorder()
		r := gin.Default()
		r.Use(func(context *gin.Context) {
			context.Set("id", uint(1))
		})

		r.DELETE("/v1/users/:id", mockApi.DeleteUser)
		req, err := http.NewRequest(http.MethodDelete, "/v1/users/1", nil)
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		mockApi.AssertNumberOfCalls(t, "DeleteUser", 1)
	})
}
