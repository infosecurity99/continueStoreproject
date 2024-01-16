package postgres

import (
	"database/sql"
	"fmt"
	"test/models"
	"test/storage"

	"github.com/google/uuid"
)

type basketProductRepo struct {
	db *sql.DB
}

func NewBasketProductRepo(db *sql.DB) storage.IBasketProductsStorage {
	return basketProductRepo{
		db: db,
	}
}

func (b basketProductRepo) Create(createBasketProduct models.CreateBasketProduct) (models.BasketProduct, error) {
	basketProductID := uuid.New()
	if _, err := b.db.Exec(`
		INSERT INTO basket_products (id, basket_id, product_id, quantity) VALUES ($1, $2, $3, $4)
	`, basketProductID, createBasketProduct.BasketID, createBasketProduct.ProductID, createBasketProduct.Quantity); err != nil {
		fmt.Println("Error while inserting basket product data", err.Error())
		return models.BasketProduct{}, err
	}
	return models.BasketProduct{
		ID:        basketProductID,
		BasketID:  createBasketProduct.BasketID,
		ProductID: createBasketProduct.ProductID,
		Quantity:  createBasketProduct.Quantity,
	}, nil
}

func (b basketProductRepo) GetByID(primaryKeysBasketProducts models.PrimaryKeysBasketProducts) (models.BasketProduct, error) {
	id := primaryKeysBasketProducts.ID
	row := b.db.QueryRow(`
		SELECT * FROM basket_products WHERE id = $1
	`, id)

	basketProduct := models.BasketProduct{}
	if err := row.Scan(&basketProduct.ID, &basketProduct.BasketID, &basketProduct.ProductID, &basketProduct.Quantity); err != nil {
		fmt.Println("Error while getting basket product data by ID", err.Error())
		return models.BasketProduct{}, err
	}

	return basketProduct, nil
}

func (b basketProductRepo) GetList(request models.GetListRequestBasketProducts) (models.BasketProductsResponse, error) {

	return models.BasketProductsResponse{}, nil
}

func (b basketProductRepo) Update(updateBasketProducts models.BasketProduct) (models.BasketProduct, error) {
	_, err := b.db.Exec(`
		UPDATE basket_products
		SET quantity = $2
		WHERE id = $1
	`, updateBasketProducts.ID, updateBasketProducts.Quantity)

	if err != nil {
		fmt.Println("Error while updating basket product data", err.Error())
		return models.BasketProduct{}, err
	}

	return updateBasketProducts, nil
}

func (b basketProductRepo) Delete(primaryKeysBasketProducts models.PrimaryKeysBasketProducts) error {
	id := primaryKeysBasketProducts.ID
	_, err := b.db.Exec(`
		DELETE FROM basket_products WHERE id = $1
	`, id)
	if err != nil {
		fmt.Println("Error while deleting basket product data", err.Error())
		return err
	}

	return nil
}
