// Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// License: MIT http://opensource.org/licenses/MIT
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vedant11/product-api/data"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewPH(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

// swagger:route GET /products products listproducts
// Returns a list of Products
// responses:
//
//	200: productsResponse
//
// GetProducts returns the product from the data store
func (p *ProductsHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (p *ProductsHandler) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse the request", http.StatusBadRequest)
	}
	data.AddProduct(&prod)
}

// swagger:route PUT /products/{id} products updateProduct
// Changes a Product
// responses:
//
//	201: updateResponse
//
// UpdateProducts returns the product from the data store
func (p *ProductsHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	id_int, _ := strconv.Atoi(id)
	prod := data.Product{ID: id_int}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse the request", http.StatusBadRequest)
	}
	err = data.UpdateProduct(&prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
