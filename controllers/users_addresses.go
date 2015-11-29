package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AddressList lists all products
func (uc UsersController) AddressList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "User.AddressList")
}

// CreateAddress - create new user address
func (uc UsersController) CreateAddress(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "User.CreateAddress")
}
