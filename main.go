// The main entrypoint for the Application

package main

import (
	"fmt"

	"github.com/FrancisVillarba/AWS-Instance-Controller/internal/server"
)

func main() {
	fmt.Println("Hello World~!")
	fmt.Println("Main: Kickstarting Application")

	apiServer := new(server.Server)
	apiServer.Init()
}
