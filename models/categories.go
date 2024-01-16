package models


import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

type UpdateCategory struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CategoriesResponse struct {
	Categories []Category `json:"categories"`
	Count      int        `json:"count"`
}

type PrimaryKeys struct {
	ID string `json:"id"`
}

type GetListRequests struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}