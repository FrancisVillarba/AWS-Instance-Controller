// General Config Loader for the AWS Instance Controller
// These Configurations are Specific for the application itself.

package config

import "os"

func InitServerConfig() {

	// Determine the environment we are running on
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "local"
	}

	// Local Environment
	if environment == "local" {
		initLocalServerConfig()
	}

	// TODO - Other Environments Setup
}

func initLocalServerConfig() {

	// TODO - Setup Any Application Specific Configs Here
}
