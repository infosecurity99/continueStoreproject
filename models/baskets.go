package models

import "github.com/google/uuid"

type Basket struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	TotalSum   int       `json:"total_sum"`
}

type CreateBasket struct {
	CustomerID uuid.UUID `json:"customer_id"`
	TotalSum   int       `json:"total_sum"`
}

type BasketsResponse struct {
	Baskets []Basket `json:"baskets"`
	Count   int      `json:"count"`
}

type PrimaryKeysBaskets struct {
	ID string `json:"id"`
}

type GetListRequestBaskets struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
