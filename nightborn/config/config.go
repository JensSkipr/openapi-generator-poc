package config

import "os"

type Config struct {
	Environment Environment
	Version     *string
	Database    map[string]string
	Sentry      map[string]string
}

// Verifies that the configuration is injected and the configuration file exists
func Initialise() Config {
	return Config{
		Environment: getEnvironment(os.Getenv("ENVIRONMENT")),
		Version:     getVersion(os.Getenv("VERSION")),
		Database: map[string]string{
			"ConnectionString": os.Getenv("DATABASE_URL"),
		},
		Sentry: map[string]string{
			"ConfigurationUrl": os.Getenv("SENTRY_URL"),
		},
	}
}
