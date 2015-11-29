package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UsersController is used for user records
type (
	UsersController struct{}
)

// BuildUsersController is ctor of UsersController
func BuildUsersController() *UsersController {
	return &UsersController{}
}

// List lists all products
func (uc UsersController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "UsersList")
}

/* session related functions */
// Signin - signin existing user
func (uc UsersController) Signin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "UserSignin")
}

// Signup - signup new user
func (uc UsersController) Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "UserSignup")
}

// Signout - sign out existing login session
func (uc UsersController) Signout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "UserSignout")
}

/* password related functions */
// ResetPassword - reset user password
func (uc UsersController) ResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "User.ResetPassword")
}

func (uc UsersController) ChangePassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "User.ChangePassword")
}

/* User info related */
// UserInfo - user info json
func (uc UsersController) UserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "User.UserInfo")
}

// UpdateUserInfo - updateuser info json
func (uc UsersController) UpdateUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "User.UpdateUserInfo")
}

/* Cart related */
func (uc UsersController) Cart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "User.Cart")
}

func (uc UsersController) UpdateCart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "User.UpdateCart")
}

func (uc UsersController) AddToCart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "User.AddToCart")
}
