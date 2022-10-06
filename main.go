package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nightborn-be/blink/skipr-test/app/controllers"
	"github.com/nightborn-be/blink/skipr-test/app/crons"
	"github.com/nightborn-be/blink/skipr-test/app/database"
	"github.com/nightborn-be/blink/skipr-test/app/gateways"
	"github.com/nightborn-be/blink/skipr-test/app/middlewares"
	"github.com/nightborn-be/blink/skipr-test/app/repositories"
	"github.com/nightborn-be/blink/skipr-test/app/routers"
	"github.com/nightborn-be/blink/skipr-test/app/usecases"
	"github.com/nightborn-be/blink/skipr-test/app/utils"
	"github.com/nightborn-be/blink/skipr-test/config"
	"github.com/nightborn-be/blink/skipr-test/config/sentry"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Load env
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return err
	}

	// Creates the config
	configs := config.Initialise()

	// Config sentry
	if err := sentry.InitialiseSentry(configs.Sentry["ConfigurationUrl"], configs.Environment, configs.Version); err != nil {
		return err
	}

	// Initialise faker generators
	utils.InitialiseFakerGenerators()

	// Initialises the router
	engine := gin.Default()

	// Creates the gateway container
	gateway := gateways.Initialise()

	// Connects to the database
	database, err := database.Initialise(configs.Database)
	if err != nil {
		return err
	}

	// Creates the repository container
	repository := repositories.Initialise(database)

	// Creates the usecase container
	usecase := usecases.Initialise(&repository, &gateway)

	// Creates the controller container
	controller := controllers.Initialise(usecase)

	// Creates middlewares
	middleware := middlewares.Initialise()

	// Creates cron-jobs
	cron := crons.Initialise(usecase)

	// Creates the routes container
	router := routers.Initialise(engine, middleware.AuthMiddleware, controller)

	// Start the crons
	if err := cron.Run(); err != nil {
		return fmt.Errorf("crons: %s", err)
	}

	// Start the router
	if err := router.Run(); err != nil {
		return fmt.Errorf("router: %s", err)
	}

	return nil
}
