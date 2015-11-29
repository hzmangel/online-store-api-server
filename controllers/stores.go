package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hzmangel/online-store-api-server/models"
	"github.com/julienschmidt/httprouter"
)

// StoresController - controller code for stores
type (
	StoresController struct{}
)

// BuildStoresController is ctor of StoresController
func BuildStoresController() *StoresController {
	return &StoresController{}
}

// List lists all stores
func (store_c StoresController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	store := models.Store{
		Name: "Greate Company",
		ID:   "FakeId",
	}

	storeJSON := store.ToJSON()
	store_c.renderJSON(w, storeJSON, 200)
}

// Create - create store record
func (store_c StoresController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	store := models.Store{}

	json.NewDecoder(r.Body).Decode(&store)

	store.ID = "FakeId"

	storeJSON := store.ToJSON()
	store_c.renderJSON(w, storeJSON, 201)
}

// Destroy - destroy store record
func (store_c StoresController) Destroy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(204)
}

func (store_c StoresController) renderJSON(w http.ResponseWriter, json []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%s", json)
}
