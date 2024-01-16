package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/models"
)

func (c Controller) Baskets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateBasket(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetBasketByID(w, r)
		} else {

		}
	case http.MethodDelete:
		c.DeleteBasket(w, r)
	case http.MethodPut:

	default:
		fmt.Println("Invalid HTTP method")
	}
}

func (c Controller) CreateBasket(w http.ResponseWriter, r *http.Request) {
	createBasket := models.CreateBasket{}

	if err := json.NewDecoder(r.Body).Decode(&createBasket); err != nil {
		fmt.Println("Error while reading data from client:", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	basket, err := c.Storage.Basket().Create(createBasket)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, basket)
}

func (c Controller) GetBasketByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	basket, err := c.Storage.Basket().GetByID(models.PrimaryKeysBaskets{ID: id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	js, err := json.Marshal(basket)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (c Controller) DeleteBasket(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Storage.Basket().Delete(models.PrimaryKeysBaskets{ID: id}); err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Basket deleted: " + id))
}
