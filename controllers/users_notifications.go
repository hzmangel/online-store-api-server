package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NotificationList - list all notifications
func (uc UsersController) NotificationList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "User.NotificationList")
}

// ReadNotification - mark given notifications as read
func (uc UsersController) ReadNotification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "User.ReadNotification")
}

//  DeleteNotification - mark given notifications as deleted (not delete it really from DB)
func (uc UsersController) DeleteNotification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "User.DeleteNotification")
}
