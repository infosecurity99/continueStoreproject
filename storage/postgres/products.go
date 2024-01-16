package postgres

import (
	"database/sql"
	"test/models"
	"test/storage"

	"github.com/google/uuid"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) storage.IProductsStorage {
	return &productRepo{
		db: db,
	}
}

func (p productRepo) Create(createProduct models.CreateProduct) (models.Product, error) {
	productID := uuid.New()

	if _, err := p.db.Exec(`
		INSERT INTO products (id, name, price, original_price, quantity, category_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, productID, createProduct.Name, createProduct.Price, createProduct.OriginalPrice, createProduct.Quantity, createProduct.CategoryID); err != nil {
		return models.Product{}, err
	}

	return models.Product{
		ID:            productID,
		Name:          createProduct.Name,
		Price:         createProduct.Price,
		OriginalPrice: createProduct.OriginalPrice,
		Quantity:      createProduct.Quantity,
		CategoryID:    createProduct.CategoryID,
	}, nil
}

func (p productRepo) GetByID(primaryKeysProducts models.PrimaryKeysProducts) (models.Product, error) {
	id := primaryKeysProducts.ID
	row := p.db.QueryRow(`
		SELECT * FROM products WHERE id = $1
	`, id)

	product := models.Product{}
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.OriginalPrice, &product.Quantity, &product.CategoryID); err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p productRepo) GetList(request models.GetListRequestProducts) (models.ProductsResponse, error) {

	return models.ProductsResponse{}, nil
}

func (p productRepo) Update(updateProduct models.UpdateProduct) (models.Product, error) {

	return models.Product{}, nil
}

func (p productRepo) Delete(primaryKeysProducts models.PrimaryKeysProducts) error {
	id := primaryKeysProducts.ID
	_, err := p.db.Exec(`
		DELETE FROM products WHERE id = $1
	`, id)
	return err
}
