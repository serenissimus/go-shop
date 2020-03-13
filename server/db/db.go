package db

import (
	"fmt"

	"go-shop/config"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewPostgres(cfg *config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DB.Postgres.Host,
		cfg.DB.Postgres.Port,
		cfg.DB.Postgres.User,
		cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Database,
	)
	if !cfg.DB.Postgres.SSL {
		connectionString += " sslmode=disable"
	}
	log.Infof("Connection string '%s'\n", connectionString)
	return gorm.Open("postgres", connectionString)

}
