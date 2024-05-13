// CSRF Middleware

package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DefaultCSRFConfig = middleware.CSRFConfig{
	Skipper:      middleware.DefaultSkipper,
	TokenLength:  32,
	TokenLookup:  "header:" + echo.HeaderXCSRFToken,
	ContextKey:   "csrf",
	CookieName:   "_csrf",
	CookieMaxAge: 86400,
}

func SetupCSRF(app *echo.Echo) {

	app.Use(middleware.CSRFWithConfig(DefaultCSRFConfig))
}
