package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

var Products = []Product{
	{Id: 1, Name: "Monster", Price: 25.0, Description: "Abra a7"},
	{Id: 2, Name: "Mac", Price: 55.0, Description: "M1"},
	{Id: 3, Name: "Lenovo", Price: 15.0, Description: "Gaming"},
	{Id: 4, Name: "Toshiba", Price: 98.0, Description: "Pro"},
}
