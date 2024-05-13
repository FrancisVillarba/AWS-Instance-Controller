// Timeout Middleware

package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DefaultTimeoutConfig = middleware.TimeoutConfig{
	Skipper:      middleware.DefaultSkipper,
	Timeout:      60 * time.Second,
	ErrorMessage: "Request timeout",
}

func SetupTimeout(app *echo.Echo) {

	app.Use(middleware.TimeoutWithConfig(DefaultTimeoutConfig))
}
