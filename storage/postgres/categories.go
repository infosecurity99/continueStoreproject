package postgres

import (
	"database/sql"
	"fmt"
	"test/models"
	"test/storage"

	"github.com/google/uuid"
)

type categoryRepo struct {
	db *sql.DB
}

// NewCategoryRepo creates a new PostgreSQL category repository.
func NewCategoryRepo(db *sql.DB) storage.ICategoriesStorage {
	return categoryRepo{
		db: db,
	}
}


func (c categoryRepo) Create(createCategory models.CreateCategory) (models.Category, error) {
	categoryID := uuid.New()

	if _, err := c.db.Exec(`
		INSERT INTO categories (id, name) VALUES ($1, $2)
	`, categoryID, createCategory.Name); err != nil {
		fmt.Println("Error while inserting category data", err.Error())
		return models.Category{}, err
	}

	return models.Category{
		ID:   categoryID,
		Name: createCategory.Name,
	}, nil
}

func (c categoryRepo) GetByIDs(rex models.PrimaryKeys) (models.Category, error) {
	id := rex.ID
	row := c.db.QueryRow(`
		SELECT * FROM categories WHERE id = $1
	`, id)

	category := models.Category{}
	if err := row.Scan(&category.ID, &category.Name); err != nil {
		fmt.Println("Error while getting category data by ID", err.Error())
		return models.Category{}, err
	}

	return category, nil
}


func (c categoryRepo) GetList(request models.GetListRequest) (models.CategoriesResponse, error) {

	return models.CategoriesResponse{}, nil
}


func (c categoryRepo) Update(updateCategory models.UpdateCategory) (models.Category, error) {

	return models.Category{}, nil
}


func (c categoryRepo) Deletes(rex models.PrimaryKeys) error {
	id := rex.ID
	_, err := c.db.Exec(`
		DELETE FROM categories WHERE id = $1
	`, id)
	if err != nil {
		fmt.Println("Error while deleting category data", err.Error())
		return err
	}

	return nil
}
