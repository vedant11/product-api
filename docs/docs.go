package docs

import "github.com/vedant11/product-api/data"

// A list of products returns in the response
// swagger:response products
type productsResponse struct {
	// ALl products in the system
	// in:body
	Body []data.Product
}

// swagger:parameters updateProduct
type updateResponse struct {
	// ALl products in the system
	// in:path
	// required:true
	Body []data.Product
}
