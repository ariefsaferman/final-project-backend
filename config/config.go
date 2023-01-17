package config

import "os"

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var (
	ENV      = getENV("ENV", "development-local")
	AppName  = getENV("APP_NAME", "go-rest-api")
	DBConfig = dbConfig{
		Host:     getENV("DB_HOST", ""),
		User:     getENV("DB_USER", ""),
		Password: getENV("DB_PASSWORD", ""),
		DBName:   getENV("DB_NAME", ""),
		Port:     getENV("DB_PORT", ""),
	}
)