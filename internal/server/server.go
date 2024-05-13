// The API Server for the Instance Controller
// Implemented using Echo

package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/FrancisVillarba/AWS-Instance-Controller/internal/middleware"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/labstack/echo/v4"
)

type Server struct {
	App *echo.Echo  `json:"App"`
	Cfg *aws.Config `json:"AWS_Config"`
}

// Initialise the Server
func (server *Server) Init() {

	// Setup the Server
	server.App = echo.New()

	// Setup waitgroup for helper
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)

	// Pass to helper to complete server init
	go server.echoInit(wg)
}

// Helper Function for Initialising Server
// Tasks that may take a while for init will be handled here instead
func (server *Server) echoInit(wg *sync.WaitGroup) {

	defer wg.Done()

	// Setup Middleware
	middleware.SetupLogger(server.App)
	middleware.SetupCORS(server.App)
	middleware.SetupCSRF(server.App)
	middleware.SetupHeaders(server.App)
	middleware.SetupLimiter(server.App)

	// Register Routes

	// Start the Server
	if err := server.App.Start(":80"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
