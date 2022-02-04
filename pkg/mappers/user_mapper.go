package mappers

import (
	"github.com/berkayersoyy/go-products-example/pkg/dto"
	"github.com/berkayersoyy/go-products-example/pkg/models"
)

func ToUser(userDTO dto.UserDTO) models.User {

	return models.User{Username: userDTO.Username, Password: userDTO.Password}
}

func ToUserDTO(user models.User) dto.UserDTO {
	return dto.UserDTO{ID: user.ID, Username: user.Password, Password: user.Password}
}

func ToUserDTOs(users []models.User) []dto.UserDTO {
	userdtos := make([]dto.UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}
