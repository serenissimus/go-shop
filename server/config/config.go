package config

type Config struct {
	DB  DBConfig
	Web WebConfig
}

func DefaultConfig() (*Config, error) {
	cfg := &Config{
		DB: DBConfig{
			Postgres: PostgresConfig{
				Host:     "localhost",
				Port:     5432,
				User:     "dev",
				Password: "secret",
				Database: "shop",
				SSL:      false,
			},
		},
		Web: WebConfig{
			Port: 8080,
			Cors: true,
		},
	}
	return cfg, nil
}
