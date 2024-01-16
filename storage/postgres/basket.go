package postgres

import (
	"database/sql"
	"fmt"
	"test/models"
	"test/storage"

	"github.com/google/uuid"
)

type basketRepo struct {
	db *sql.DB
}


func NewBasketRepo(db *sql.DB) storage.IBasketsStorage {
	return &basketRepo{
		db: db,
	}
}

func (b *basketRepo) Create(createBasket models.CreateBasket) (models.Basket, error) {
	basketID := uuid.New()
	if _, err := b.db.Exec(`
		INSERT INTO baskets (id, customer_id, total_sum) VALUES ($1, $2, $3)
	`, basketID, createBasket.CustomerID, createBasket.TotalSum); err != nil {
		fmt.Println("Error while inserting basket data", err.Error())
		return models.Basket{}, err
	}

	return models.Basket{
		ID:         basketID,
		CustomerID: createBasket.CustomerID,
		TotalSum:   createBasket.TotalSum,
	}, nil
}

func (b *basketRepo) GetByID(primaryKeysBaskets models.PrimaryKeysBaskets) (models.Basket, error) {
	id := primaryKeysBaskets.ID
	row := b.db.QueryRow(`
		SELECT * FROM baskets WHERE id = $1
	`, id)

	basket := models.Basket{}
	if err := row.Scan(&basket.ID, &basket.CustomerID, &basket.TotalSum); err != nil {
		fmt.Println("Error while getting basket data by ID", err.Error())
		return models.Basket{}, err
	}

	return basket, nil
}

func (b *basketRepo) GetList(request models.GetListRequestBaskets) (models.BasketsResponse, error) {
	return models.BasketsResponse{}, nil
}

func (b *basketRepo) Update(updateBasket models.Basket) (models.Basket, error) {
	return models.Basket{}, nil
}

func (b *basketRepo) Delete(primaryKeysBaskets models.PrimaryKeysBaskets) error {
	id := primaryKeysBaskets.ID
	_, err := b.db.Exec(`
		DELETE FROM baskets WHERE id = $1
	`, id)
	if err != nil {
		fmt.Println("Error while deleting basket data", err.Error())
		return err
	}

	return nil
}
