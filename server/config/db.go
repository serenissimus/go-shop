package config

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSL      bool
}

type DBConfig struct {
	Postgres PostgresConfig
}
