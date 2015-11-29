package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	StoresController struct{}
)

// BuildStoresController is ctor of StoresController
func BuildStoresController() *StoresController {
	return &StoresController{}
}

// List lists all stores
func (store_c StoresController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "StoreList")
}
