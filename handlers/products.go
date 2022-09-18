// Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"log"
	"net/http"

	"github.com/vedant11/product-api/data"
)

// swagger: route GET /products products listProducts
// Returns a list of products
// responses:
// 	200:productResponse

type ProductsHandler struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}
func (p *ProductsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusNotImplemented)
}

// returns the products from data store
func (p *ProductsHandler) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	// JSON encoding instead of Marshal
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (p *ProductsHandler) addProduct(rw http.ResponseWriter, r *http.Request) {
	prod := data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse the request", http.StatusBadRequest)
	}
	data.AddProduct(&prod)
}
