package config

type Environment string

const (
	ENVIRONMENT_PRODUCTION Environment = "production"
	ENVIRONMENT_STAGING    Environment = "staging"
	ENVIRONMENT_LOCAL      Environment = "local"
)

var Environments = map[string]Environment{
	string(ENVIRONMENT_PRODUCTION): ENVIRONMENT_PRODUCTION,
	string(ENVIRONMENT_STAGING):    ENVIRONMENT_STAGING,
	string(ENVIRONMENT_LOCAL):      ENVIRONMENT_LOCAL,
}

func getEnvironment(env string) Environment {
	environment, ok := Environments[env]
	if !ok {
		environment = ENVIRONMENT_LOCAL
	}
	return environment
}
