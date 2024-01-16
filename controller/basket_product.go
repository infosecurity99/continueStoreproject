package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/models"
)

func (c Controller) BasketProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateBasketProduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetBasketProductByID(w, r)
		} else {

		}
	case http.MethodDelete:
		c.DeleteBasketProduct(w, r)
	case http.MethodPut:

	default:
		fmt.Println("Invalid HTTP method")
	}
}

func (c Controller) CreateBasketProduct(w http.ResponseWriter, r *http.Request) {
	createBasketProduct := models.CreateBasketProduct{}

	if err := json.NewDecoder(r.Body).Decode(&createBasketProduct); err != nil {
		fmt.Println("Error while reading data from client:", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	basketProduct, err := c.Storage.BasketProduct().Create(createBasketProduct)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, basketProduct)
}

func (c Controller) GetBasketProductByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	basketProduct, err := c.Storage.BasketProduct().GetByID(models.PrimaryKeysBasketProducts{ID: id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	js, err := json.Marshal(basketProduct)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (c Controller) DeleteBasketProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Storage.BasketProduct().Delete(models.PrimaryKeysBasketProducts{ID: id}); err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Basket Product deleted: " + id))
}


