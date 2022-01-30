package dto

type ProductDTO struct {
	ID          uint    `json:"id,string,omitempty"`
	Name        string  `json:"name" validate:"required,min=2,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
