package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for a Product API
// swagger:model
type Product struct {
	// the id of the user
	// required:true
	// min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"SKU"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *Product) FromJSON(data io.Reader) error {
	e := json.NewDecoder(data)
	return e.Decode(p)
}
func GetProducts() Products {
	return productList
}
func AddProduct(p *Product) {
	productList = append(productList, p)
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func UpdateProduct(p *Product) error {
	prod, index, err := findAndUpdateProduct(p.ID)
	fmt.Println(index, prod)
	return err

}
func findAndUpdateProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
