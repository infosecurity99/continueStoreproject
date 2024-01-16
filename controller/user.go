package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/models"
)

func (c Controller) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateUser(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByID(w, r)
		} else {

		}
	case http.MethodDelete:
		c.Delete(w, r)
	case http.MethodPut:

	default:
		fmt.Println("this is not case ")
	}
}

//create
func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := models.CreateUser{}

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := c.Storage.User().Create(createUser)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusCreated, user)
}

//getbyid
func (c Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	users, err := c.Storage.User().GetByID(models.PrimaryKey{ID: id})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	js, err := json.Marshal(users)
	if err != nil {
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

//delete
func (c Controller) Delete(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Storage.User().Delete(models.PrimaryKey{ID: id}); err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`this is delete` + id))
}

