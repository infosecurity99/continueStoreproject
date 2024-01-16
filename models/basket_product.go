package models

import "github.com/google/uuid"

type BasketProduct struct {
	ID        uuid.UUID `json:"id"`
	BasketID  uuid.UUID `json:"basket_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type CreateBasketProduct struct {
	BasketID  uuid.UUID `json:"basket_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type PrimaryKeysBasketProducts struct {
	ID string `json:"id"`
}

type GetListRequestBasketProducts struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type BasketProductsResponse struct {
	BasketProducts []BasketProduct `json:"basket_products"`
	Count          int             `json:"count"`
}
