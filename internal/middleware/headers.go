// Security Headers Middleware

package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DefaultHeadersConfig = middleware.SecureConfig{
	Skipper:               middleware.DefaultSkipper,
	XSSProtection:         "1; mode=block",
	ContentTypeNosniff:    "nosniff",
	XFrameOptions:         "SAMEORIGIN",
	HSTSMaxAge:            3600,
	ContentSecurityPolicy: "default-src 'self'",
}

func SetupHeaders(app *echo.Echo) {

	app.Use(middleware.SecureWithConfig(DefaultHeadersConfig))
}
