package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/models"
)

func (c Controller) Categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateCategory(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetCategoryByID(w, r)
		} else {
	
		}
	case http.MethodDelete:
		c.DeleteCategory(w, r)
	case http.MethodPut:
	
	default:
		fmt.Println("Invalid HTTP method")
	}
}


func (c Controller) CreateCategory(w http.ResponseWriter, r *http.Request) {
	createCategory := models.CreateCategory{}

	if err := json.NewDecoder(r.Body).Decode(&createCategory); err != nil {
		fmt.Println("Error while reading data from client:", err.Error())
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	category, err := c.Storage.Categories().Create(createCategory)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusCreated, category)
}

func (c Controller) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	category, err := c.Storage.Categories().GetByIDs(models.PrimaryKeys{ID: id})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	js, err := json.Marshal(category)
	if err != nil {
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (c Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Storage.Categories().Deletes(models.PrimaryKeys{ID: id}); err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Category deleted: " + id))
}
