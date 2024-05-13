// Config Loader for AWS Instance Controller

package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

func InitConfig() aws.Config {

	// Determine the environment we are running on
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "local"
	}

	// Local Environment
	if environment == "local" {
		config := initLocalAWSConfig()
		return config
	}

	// Deployed within aws (i.e. ec2, lambda with appropriate IAM roles set)
	if environment == "aws" {
		config := initAwsConfig()
		return config
	}

	// Deployed to a third party server / service
	if environment == "external" {
		config := initExternalAWSConfig()
		return config
	}

	panic("Unknown Environment Set, Only 'local', 'aws' & 'external' is supported.")
}

// Helper Function - Initialise Configuration if Local
func initLocalAWSConfig() aws.Config {

	hasError := false

	cfg, err := godotenv.Read(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
		hasError = true
	}

	// Check if the required number of environment variables exist

	_, present := cfg["AWS_REGION"]
	if !present {
		log.Fatal("Missing Environment Variable 'AWS_REGION' in .env.local")
		hasError = true
	}

	_, present = cfg["AWS_ACCESS_KEY"]
	if !present {
		log.Fatal("Missing Environment Variable 'AWS_ACCESS_KEY' in .env.local")
		hasError = true
	}

	_, present = cfg["AWS_SECRET_KEY"]
	if !present {
		log.Fatal("Missing Environment Variable 'AWS_SECRET_KEY' in .env.local")
		hasError = true
	}

	if hasError {
		panic("Missing Environment Variables required to Run this Application. Please check the logs for more information")
	}

	// Load the configuration & set it on the system
	godotenv.Load(".env.local")

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg["AWS_REGION"]))
	if err != nil {
		log.Fatalf("Failed to initialise AWS SDK v2 Configuration, %v", err)
	}

	return config
}

// Helper Function - Initialise Configuration if Deployed within AWS
func initAwsConfig() aws.Config {

	// Check if environment variables exist, we should expect these when in AWS
	hasError := false

	_, present := os.LookupEnv("AWS_DEFAULT_REGION")
	if !present {
		log.Fatal("Missing Expected Environment Variable 'AWS_DEFAULT_REGION'")
		hasError = true
	}

	reg, present := os.LookupEnv("AWS_REGION")
	if !present {
		log.Fatal("Missing Expected Environment Variable 'AWS_REGION'")
		hasError = true
	}

	if hasError {
		panic("Missing Environment Variables required to Run this Application. Please check the logs for more information")
	}

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(reg))
	if err != nil {
		log.Fatalf("Failed to initialise AWS SDK v2 Configuration, %v", err)
	}

	return config
}

// Helper Function - Initialise Configuration if Deployed to a Third Party Service
// This assumes it is running within Docker and the environment variables are passed thru / mounted on Runtime.
func initExternalAWSConfig() aws.Config {

	// Check if environment variables exist
	hasError := false

	_, present := os.LookupEnv("AWS_REGION")
	if !present {
		log.Fatal("Missing Expected Environment Variable 'AWS_REGION'")
		hasError = true
	}

	_, present = os.LookupEnv("AWS_REGION")
	if !present {
		log.Fatal("Missing Expected Environment Variable 'AWS_REGION'")
		hasError = true
	}

	_, present = os.LookupEnv("AWS_ACCESS_KEY")
	if !present {
		log.Fatal("Missing Expected Environment Variable 'AWS_ACCESS_KEY'")
		hasError = true
	}

	_, present = os.LookupEnv("AWS_SCRET_KEY")
	if !present {
		log.Fatal("Missing Expected Environment Variable 'AWS_SECRET_KEY'")
		hasError = true
	}

	if hasError {
		panic("Missing Environment Variables required to Run this Application. Please check the logs for more information")
	}

	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to initialise AWS SDK v2 Configuration, %v", err)
	}

	return config
}
