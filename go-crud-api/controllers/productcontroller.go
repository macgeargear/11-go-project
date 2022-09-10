package controllers

import (
	"encoding/json"
	"go-crud-api/database"
	"go-crud-api/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product) // return the newly created product
}

func checkIfProdExist(productId string) bool {
	var product entities.Product
	database.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}

func GetProductsById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProdExist(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []entities.Product
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProdExist(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&productId)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
