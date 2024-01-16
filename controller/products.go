package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/models"
)

func (c Controller) Products(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateProduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetProductByID(w, r)
		} else {
			c.GetProductsList(w, r)
		}
	case http.MethodDelete:
		c.DeleteProduct(w, r)
	case http.MethodPut:
		c.UpdateProduct(w, r)
	default:
		fmt.Println("Invalid HTTP method")
	}
}


func (c Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	createProduct := models.CreateProduct{}

	if err := json.NewDecoder(r.Body).Decode(&createProduct); err != nil {
		fmt.Println("Error while reading data from client", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	product, err := c.Storage.Products().Create(createProduct)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, product)
}

func (c Controller) GetProductByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	product, err := c.Storage.Products().GetByID(models.PrimaryKeysProducts{ID: id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	js, err := json.Marshal(product)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (c Controller) GetProductsList(w http.ResponseWriter, r *http.Request) {
	
	handleResponse(w, http.StatusOK, "List of products")
}

func (c Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	updateProduct := models.UpdateProduct{}

	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil {
		fmt.Println("Error while reading data from client", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	updatedProduct, err := c.Storage.Products().Update(updateProduct)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, updatedProduct)
}


func (c Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Storage.Products().Delete(models.PrimaryKeysProducts{ID: id}); err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, fmt.Sprintf("Product with ID %s has been deleted", id))
}


func handleResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}


	w.WriteHeader(statusCode)
	w.Write(js)
}
