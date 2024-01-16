package models

import "github.com/google/uuid"

type Product struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	OriginalPrice int       `json:"original_price"`
	Quantity      int       `json:"quantity"`
	CategoryID    uuid.UUID `json:"category_id"`
}

type CreateProduct struct {
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	OriginalPrice int       `json:"original_price"`
	Quantity      int       `json:"quantity"`
	CategoryID    uuid.UUID `json:"category_id"`
}

type UpdateProduct struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

type ProductsResponse struct {
	Products []Product `json:"products"`
	Count    int       `json:"count"`
}

type PrimaryKeysProducts struct {
	ID string `json:"id"`
}

type GetListRequestProducts struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
