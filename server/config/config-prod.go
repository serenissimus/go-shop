package config

import "time"

func ProdConfig() (*Config, error) {
	cfg := &Config{
		DB: DBConfig{
			Postgres: PostgresConfig{
				Host:     "db",
				Port:     5432,
				User:     "dev",
				Password: "secret",
				Database: "shop",
				SSL:      false,
			},
			MaxTries: 10,
			Timeout:  1 * time.Second,
		},
		Web: WebConfig{
			Port: 8080,
			Cors: false,
		},
	}
	return cfg, nil
}
