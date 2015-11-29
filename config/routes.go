package config

import (
	"github.com/hzmangel/online-store-api-server/controllers"
	"github.com/julienschmidt/httprouter"
)

// Routes defines all routes used in the project
func Routes() *httprouter.Router {
	r := httprouter.New()

	storeCtrl := controllers.BuildStoresController()
	prodCtrl := controllers.BuildProductsController()
	prodCategoryCtrl := controllers.BuildProductCategoriesController()
	orderCtrl := controllers.BuildOrdersController()
	userCtrl := controllers.BuildUsersController()

	//////////////////////////////
	// User related routes
	//////////////////////////////
	r.POST("/auth/signup", userCtrl.Signup)
	r.POST("/auth/signin", userCtrl.Signin)
	r.DELETE("/auth/signout", userCtrl.Signout)

	r.POST("/password/reset", userCtrl.ResetPassword)
	r.PATCH("/password/change", userCtrl.ChangePassword)

	r.GET("/me", userCtrl.UserInfo)
	r.PATCH("/me", userCtrl.UpdateUserInfo)

	r.GET("/cart", userCtrl.Cart)
	r.PATCH("/cart", userCtrl.UpdateCart)
	r.PATCH("/add_to_cat", userCtrl.AddToCart)

	r.GET("/addresses", userCtrl.AddressList)
	r.POST("/addresses", userCtrl.CreateAddress)

	r.GET("/notifications", userCtrl.NotificationList)
	r.PATCH("/notifications/read", userCtrl.ReadNotification)
	r.DELETE("/notifications/delete", userCtrl.DeleteNotification)

	//////////////////////////////
	// Store related routes
	//////////////////////////////
	r.GET("/stores/list", storeCtrl.List)

	//////////////////////////////
	// Product related routes
	//////////////////////////////
	r.GET("/products/list", prodCtrl.List)
	r.POST("/products/create", prodCtrl.Create)
	r.GET("/products/show", prodCtrl.Show)
	r.PATCH("/products/update", prodCtrl.Update)
	r.DELETE("/products/destroy", prodCtrl.Destroy)

	r.GET("/products/search", prodCtrl.Search)
	r.GET("/products/online", prodCtrl.Online)
	r.GET("/products/offline", prodCtrl.Offline)

	//////////////////////////////
	// Product category related routes
	//////////////////////////////
	r.GET("/product_categories/list", prodCategoryCtrl.List)
	r.POST("/product_categories/create", prodCategoryCtrl.Create)
	r.GET("/product_categories/show", prodCategoryCtrl.Show)
	r.PATCH("/product_categories/update", prodCategoryCtrl.Update)
	r.DELETE("/product_categories/destroy", prodCategoryCtrl.Destroy)

	//////////////////////////////
	// Order related routes
	//////////////////////////////
	r.GET("/orders", orderCtrl.List)
	r.GET("/products/search", orderCtrl.Search)
	r.POST("/orders", orderCtrl.Create)
	r.GET("/orders/show", orderCtrl.Show)
	r.PATCH("/orders/update", orderCtrl.Update)
	r.DELETE("/orders/destroy", orderCtrl.Destroy)

	// Order state change API
	r.POST("/orders/pay", orderCtrl.Pay)
	r.POST("/orders/received", orderCtrl.Received)
	r.POST("/orders/refund", orderCtrl.CreateRefundRequest)
	r.PATCH("/orders/refund", orderCtrl.UpdateRefundRequest)
	r.POST("/orders/refund_express", orderCtrl.CreateRefundExpress)
	r.PATCH("/orders/efund_express", orderCtrl.UpdateRefundExpress)

	return r
}
