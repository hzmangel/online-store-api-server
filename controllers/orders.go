package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	OrdersController struct{}
)

// BuildOrdersController is ctor of OrdersController
func BuildOrdersController() *OrdersController {
	return &OrdersController{}
}

// List lists all Orders
func (pc OrdersController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "OrderList")
}

// Show shows single Order
func (pc OrdersController) Show(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "OrderShow")
}

// Create shows single Order
func (pc OrdersController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "OrderCreate")
}

// Update updates single Order record
func (pc OrdersController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "OrderUpdate")
}

// Destroy destroys Order record and associated data
func (pc OrdersController) Destroy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "OrderUpdate")
}

// Search  - search Orders
func (pc OrdersController) Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Orders.Search")
}
func (pc OrdersController) Pay(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Orders.Pay")
}
func (pc OrdersController) Received(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Orders.Received")
}
func (pc OrdersController) CreateRefundRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Orders.RefundRequest")
}
func (pc OrdersController) UpdateRefundRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Orders.UpdateRefundRequest")
}
func (pc OrdersController) CreateRefundExpress(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Orders.CreateRefundExpress")
}
func (pc OrdersController) UpdateRefundExpress(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Orders.UpdateRefundExpress")
}
