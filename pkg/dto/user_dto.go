package dto

// swagger:model UserDTO
type UserDTO struct {
	ID       uint   `json:"id,string,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
