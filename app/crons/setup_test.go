package crons

import (
	"github.com/joho/godotenv"
	"github.com/nightborn-be/blink/skipr-test/app/database"
	"github.com/nightborn-be/blink/skipr-test/app/gateways"
	"github.com/nightborn-be/blink/skipr-test/app/repositories"
	"github.com/nightborn-be/blink/skipr-test/app/usecases"
	"github.com/nightborn-be/blink/skipr-test/config"
	"github.com/samber/lo"
)

// Setup test cron
func SetupTestCron() (*CronJob, error) {
	// Load env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Creates the config
	config := config.Initialise()

	// Setup database
	db, err := database.Initialise(config.Database)
	if err != nil {
		return nil, err
	}

	// Setup in memory db for repositories & in-memory gateways
	testRepository := repositories.Initialise(db)
	testGateway, _ := gateways.InitialiseTest()

	// Setup usecases with repositories, gateways
	usecase := usecases.Initialise(&testRepository, &testGateway)

	// Setup crons
	return lo.ToPtr(Initialise(usecase)), nil
}
