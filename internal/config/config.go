package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

const defaultFetchSize = 10

var config *Config

// App is config for app
type App struct {
	Port string
	Env  string
}

type PostProvider struct {
	URL       string
	FetchSize int
}

// DB is config for db
type DB struct {
	Host     string
	Port     string
	Name     string
	Password string
	User     string
	SSL      string
	Timezone string
}

// Config is config struct
type Config struct {
	App          App
	DB           DB
	PostProvider PostProvider
}

// GetConfig returns app config
func GetConfig() (*Config, error) {
	if config == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, errors.Wrap(err, "could not load app-config")
		}

		fetcherSize, err := strconv.Atoi(os.Getenv("POST_PROVIDER_FETCH_SIZE"))
		if err != nil {
			fetcherSize = defaultFetchSize
		}

		config = &Config{
			App: App{
				Port: os.Getenv("APP_PORT"),
				Env:  os.Getenv("APP_ENV"),
			},
			DB: DB{
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Name:     os.Getenv("DB_NAME"),
				Password: os.Getenv("DB_PASSWORD"),
				User:     os.Getenv("DB_USER"),
				SSL:      os.Getenv("DB_SSL"),
				Timezone: os.Getenv("DB_TIMEZONE"),
			},
			PostProvider: PostProvider{
				URL:       os.Getenv("POST_PROVIDER_URL"),
				FetchSize: fetcherSize,
			},
		}
	}

	return config, nil
}
