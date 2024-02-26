package settings

import (
	"os"
)

type Settings struct {
	GOENV             string
	Port              string
	JWTSecret         string
	JWTExpirationDelta int
}

var env string
var settings = Settings{}

// Init reads env variables and assings settings object
func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		// fmt.Println("Warning: Setting development environment due to lack of GO_ENV value")
		env = "development"
	}
	LoadSettingsByEnv(env)

}

// LoadSettingsByEnv sets the global object
func LoadSettingsByEnv(env string) {
	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "4444"
	}
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		jwtSecret = "mysecretpw"
	}
	if env == "development" {

	} else if env == "production" {

	}
	settings = Settings{
		GOENV:             env,
		Port:              port,
		JWTSecret:         jwtSecret,
		JWTExpirationDelta: 720,
	}
}

// GetEnvironment returns env variable
func GetEnvironment() string {
	return env
}

// Get returns the settings
func Get() Settings {
	if settings == (Settings{}) {
		Init()
	}
	return settings
}

func isTestEnvironment() bool {
	return env == "tests"
}
