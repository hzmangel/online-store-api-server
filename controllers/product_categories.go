package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ProductCategoriesController is used for all categories of products in app
type (
	ProductCategoriesController struct{}
)

// BuildProductCategoriesController is ctor of ProductCategoriesController
func BuildProductCategoriesController() *ProductCategoriesController {
	return &ProductCategoriesController{}
}

// List lists all products
func (prod_cat_c ProductCategoriesController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductCategoryList")
}

func (prod_cat_c ProductCategoriesController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductCategory.Create")
}
func (prod_cat_c ProductCategoriesController) Show(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductCategory.Create")
}
func (prod_cat_c ProductCategoriesController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductCategory.Create")
}
func (prod_cat_c ProductCategoriesController) Destroy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ProductCategory.Create")
}
