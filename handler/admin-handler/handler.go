package adminHandler

import (
	"Kingcrusher9764/go-server/handler/auth-handler"
	"Kingcrusher9764/go-server/templ/admin-views"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	homeView := adminView.Home()
	var msgs []string

	return authHandler.RenderView(c, adminView.AdminIndex(
		"Home",
		true,
		msgs,
		msgs,
		homeView,
	))
}

func Customers(c echo.Context) error {
	customersView := adminView.Customers()
	var msgs []string

	return authHandler.RenderView(c, adminView.AdminIndex(
		"Customers",
		true,
		msgs,
		msgs,
		customersView,
	))
}

func Orders(c echo.Context) error {
	ordersView := adminView.Orders()
	var msgs []string

	return authHandler.RenderView(c, adminView.AdminIndex(
		"Orders",
		true,
		msgs,
		msgs,
		ordersView,
	))
}
