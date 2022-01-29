package dto

type ProductDTO struct {
	ID          uint    `json:"id,string,omitempty"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
