package config

import (
	"os"
)

const (
	LOCAL       = "local"
	DEVELOPMENT = "development"
)

// ENVIRONMENT:
const ENVIRONMENT string = LOCAL // LOCAL, DEVELOPMENT, PRODUCTION

var env = map[string]map[string]string{
	// local environment configuration
	"local": {
		"PORT":            "80",
		"POSTGRES_HOST":   "localhost",
		"POSTGRES_PORT":   "5432",
		"POSTGRES_USER":   "postgres",
		"POSTGRES_PASS":   "meliodasten10",
		"POSTGRES_SCHEMA": "user_db",

		"SECRET_KEY": "melio",
		"APP_NAME":   "backend_api",
	},

	// development environment configuration
	"development": {
		"PORT": "8080",

		"MYSQL_HOST":   "",
		"MYSQL_PORT":   "",
		"MYSQL_USER":   "",
		"MYSQL_PASS":   "",
		"MYSQL_SCHEMA": "",
	},
}

// CONFIG : global configuration
var CONFIG = env[ENVIRONMENT]

// Getenv : function for Environment Lookup
func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
