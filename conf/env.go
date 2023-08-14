package conf

import (
	"fmt"

	env "github.com/caarlos0/env"
)

type Config struct {
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"DB_PORT"`
	DBUser     string `env:"DB_USER" envDefault:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME" envDefault:"DB_NAME"`
	Port       string `env:"PORT" envDefault:"5432"`

	MaxOpenConns    int `env:"MAX_OPEN_CONNS" envDefault:"25"`
	MaxIdleConns    int `env:"MAX_IDLE_CONNS" envDefault:"25"`
	ConnMaxLifetime int `env:"CONN_MAX_LIFETIME" envDefault:"600"`
}

var cfg Config

func SetEnv() {
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("Failed to parse env %v", err)
		return
	}
}

func GetEnv() Config {
	return cfg
}
