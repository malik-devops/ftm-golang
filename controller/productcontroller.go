

package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"test/database"
	"test/entity"
	"strconv"

	"github.com/gorilla/mux"
)


//GetAllproduct get all product data
func GetAllproduct(w http.ResponseWriter, r *http.Request) {
	var products []entity.Product
	database.Connector.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

//GetproductByID returns product with specific ID
func GetproductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var product entity.Product
	database.Connector.First(&product, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

//Createproduct creates product
func Createproduct(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Product
	json.Unmarshal(requestBody, &product)

	database.Connector.Create(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

//UpdateproductByID updates product with respective ID
func UpdateproductByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Product
	json.Unmarshal(requestBody, &product)
	database.Connector.Save(&product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

//DeletproductByID delete's product with specific ID
func DeletproductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var product entity.Product
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&product)
	w.WriteHeader(http.StatusNoContent)
}