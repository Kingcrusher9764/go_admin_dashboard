package authHandler

import (
	authView "Kingcrusher9764/go-server/templ/auth-views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	loginView := authView.Login()
	var msgs []string

	if c.Request().Method == "POST" {
		// do the login procedures
	}

	return RenderView(c, authView.LoginIndex(
		"Login",
		false,
		msgs,
		msgs,
		loginView,
	))
}

func RegisterHandler(c echo.Context) error {
	registerView := authView.Register()
	var msgs []string

	if c.Request().Method == "POST" {
		// do the register procedures
	}

	return RenderView(c, authView.RegisterIndex(
		"Register",
		false,
		msgs,
		msgs,
		registerView,
	))
}

func RenderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
