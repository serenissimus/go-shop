package config

import "time"

func DevConfig() (*Config, error) {
	cfg := &Config{
		DB: DBConfig{
			Postgres: PostgresConfig{
				Host:     "0.0.0.0",
				Port:     5432,
				User:     "dev",
				Password: "secret",
				Database: "shop",
				SSL:      false,
			},
			MaxTries: 5,
			Timeout:  1 * time.Second,
		},
		Web: WebConfig{
			Port: 8080,
			Cors: true,
		},
	}
	return cfg, nil
}
