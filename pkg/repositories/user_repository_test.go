package repositories

import (
	"fmt"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestRepositoryGetUserById() {
	user := models.User{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, Username: "test-username", Password: "test-password"}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` = 1)) ORDER BY `users`.`id` ASC LIMIT 1")).
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "username", "password"}).
			AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.DeletedAt, user.Username, user.Password))
	res := s.userRepository.GetUserByID(user.ID)

	require.Equal(s.T(), user, res)
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}
}

func (s *Suite) TestRepositoryGetAllUsers() {
	user := models.User{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, Username: "test-username", Password: "test-password"}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` ASC LIMIT 1")).
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "username", "password"}).
			AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.DeletedAt, user.Username, user.Password))
	res := s.userRepository.GetAllUsers()

	require.Equal(s.T(), user, res[0])
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}
}
func (s *Suite) TestRepositoryAddUser() {
	user := models.User{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, Username: "test-username", Password: "test-password"}
	prep := s.mock.ExpectPrepare("INSERT INTO users (username, password) VALUES (?, ?)")
	prep.ExpectExec().
		WithArgs(user.Username, user.Password).
		WillReturnResult(sqlmock.NewResult(0, 1))
	res := s.userRepository.AddUser(user)

	require.Equal(s.T(), user, res)
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}
}
func (s *Suite) TestRepositoryDeleteUser() {
	user := models.User{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, Username: "test-username", Password: "test-password"}
	prep := s.mock.ExpectPrepare("DELETE from users WHERE id = ?")
	prep.ExpectExec().
		WithArgs(user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.userRepository.DeleteUser(user)
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}
}
