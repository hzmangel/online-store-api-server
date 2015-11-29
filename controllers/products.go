package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	ProductsController struct{}
)

// BuildProductsController is ctor of ProductsController
func BuildProductsController() *ProductsController {
	return &ProductsController{}
}

// List lists all products
func (pc ProductsController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductList")
}

// Show shows single product
func (pc ProductsController) Show(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductShow")
}

// Create shows single product
func (pc ProductsController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductCreate")
}

// Update updates single product record
func (pc ProductsController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductUpdate")
}

// Destroy destroys product record and associated data
func (pc ProductsController) Destroy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductUpdate")
}

// Search  - search products
func (pc ProductsController) Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Products.Search")
}

// Online  - mark products as online
func (pc ProductsController) Online(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Products.Online")
}

// Offline - mark products as offline
func (pc ProductsController) Offline(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Products.Offline")
}
