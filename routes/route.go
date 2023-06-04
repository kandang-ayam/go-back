package routes

import (
	"github.com/labstack/echo/v4"
	"point-of-sale/app/controller"
	"point-of-sale/app/controller/admin"
	"point-of-sale/app/middleware"
)

func Route(e *echo.Echo) {
	api := e.Group("api/v1")

	// Role cashier
	RouteCashier := api.Group("/cashier")
	RouteCashier.POST("/login", controller.LoginCashier)
	RouteCashier.Use(middleware.JWTMiddleware)

	// Role admin
	RouteAdmin := api.Group("/admin")
	RouteAdmin.POST("/login", controller.LoginAdmin)
	RouteAdmin.Use(middleware.JWTMiddleware, middleware.AdminMiddleware)

	// Membership endpoints
	api.GET("/admin/membership", admin.GetMembership, middleware.JWTMiddleware)
	api.POST("/admin/membership", admin.AddMembership, middleware.JWTMiddleware)
	api.POST("/admin/membership/point", admin.AddPoint, middleware.JWTMiddleware)
	api.PUT("/admin/membership/:id", admin.EditMembership, middleware.JWTMiddleware)
	api.DELETE("/admin/membership/:id", admin.DeleteMembership, middleware.JWTMiddleware)

	// Cashier endpoints
	api.GET("/admin/cashier", admin.GetCashier, middleware.JWTMiddleware)
	api.POST("/admin/cashier", admin.AddCashier, middleware.JWTMiddleware)
	api.PUT("/admin/cashier/:id", admin.EditCashier, middleware.JWTMiddleware)
	api.DELETE("/admin/cashier/:id", admin.DeleteCashier, middleware.JWTMiddleware)
}
