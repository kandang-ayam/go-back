package routes

import (
	"point-of-sale/app/controller"
	"point-of-sale/app/controller/admin"
	"point-of-sale/app/middleware"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	api := e.Group("api/v1")

	//Role cashier
	RouteCashier := api.Group("/cashier")
	RouteCashier.POST("/login", controller.LoginCashier)
	RouteCashier.Use(middleware.JWTMiddleware)
	{

	}

	RouteAdmin := api.Group("/admin")
	RouteAdmin.POST("/login", controller.LoginAdmin)
	RouteAdmin.Use(middleware.JWTMiddleware, middleware.AdminMiddleware)
	{
		RouteAdmin.GET("/cashier", admin.GetCashier)
		RouteAdmin.POST("/cashier", admin.AddCashier)
		RouteAdmin.PUT("/cashier/:id", admin.EditCashier)
		RouteAdmin.DELETE("/cashier/:id", admin.DeleteCashier)

		RouteAdmin.GET("/membership", admin.GetMembership)
		RouteAdmin.POST("/membership", admin.AddMembership)
		RouteAdmin.POST("/membership/point", admin.AddPoint)
		RouteAdmin.PUT("/membership/:id", admin.EditMembership)
		RouteAdmin.DELETE("/membership/:id", admin.DeleteMembership)
	}

}
