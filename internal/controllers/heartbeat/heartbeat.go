// API Controller for Health Check
// Created & maintained by Francis Villarba

package heartbeat

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(ctx echo.Context) error {
	fmt.Println("Heartbeat Check Hit - I'm pretty sure we are alive!")
	return ctx.String(http.StatusOK, "AWS Instance Controller -- Ready to Rock and Roll!")
}
