// CORS Middleware

package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DefaultCORSConfig = middleware.CORSConfig{

	Skipper:      middleware.DefaultSkipper,
	AllowOrigins: []string{"*"},
	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost},
}

func SetupCORS(app *echo.Echo) {

	app.Use(middleware.CORSWithConfig(DefaultCORSConfig))
}
