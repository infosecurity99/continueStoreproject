package storage

import "test/models"

type IStorage interface {
	Close()
	User() IUserStorage
	Categories() ICategoriesStorage
	Products() IProductsStorage
	Basket() IBasketsStorage
	BasketProduct() IBasketProductsStorage
}

//users  interface
type IUserStorage interface {
	Create(models.CreateUser) (models.User, error)
	GetByID(models.PrimaryKey) (models.User, error)
	GetList(models.GetListRequest) (models.UsersResponse, error)
	Update(models.UpdateUser) (models.User, error)
	Delete(models.PrimaryKey) error
}

//categoryinterface
type ICategoriesStorage interface {
	Create(models.CreateCategory) (models.Category, error)
	GetByIDs(models.PrimaryKeys) (models.Category, error)
	GetList(models.GetListRequest) (models.CategoriesResponse, error)
	Update(models.UpdateCategory) (models.Category, error)
	Deletes(models.PrimaryKeys) error
}

//productinterface
type IProductsStorage interface {
	Create(models.CreateProduct) (models.Product, error)
	GetByID(models.PrimaryKeysProducts) (models.Product, error)
	GetList(models.GetListRequestProducts) (models.ProductsResponse, error)
	Update(models.UpdateProduct) (models.Product, error)
	Delete(models.PrimaryKeysProducts) error
}

//basket interface
type IBasketsStorage interface {
	Create(models.CreateBasket) (models.Basket, error)
	GetByID(models.PrimaryKeysBaskets) (models.Basket, error)
	GetList(models.GetListRequestBaskets) (models.BasketsResponse, error)
	Update(models.Basket) (models.Basket, error)
	Delete(models.PrimaryKeysBaskets) error
}

//basket_product

type IBasketProductsStorage interface {
	Create(models.CreateBasketProduct) (models.BasketProduct, error)
	GetByID(models.PrimaryKeysBasketProducts) (models.BasketProduct, error)
	GetList(models.GetListRequestBasketProducts) (models.BasketProductsResponse, error)
	Update(models.BasketProduct) (models.BasketProduct, error)
	Delete(models.PrimaryKeysBasketProducts) error
}