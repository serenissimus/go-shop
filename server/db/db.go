package db

import (
	"fmt"
	"time"

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
	log.Infof("Connection string is '%s'", connectionString)

	var (
		db  *gorm.DB
		err error
	)

	for i := 0; i < cfg.DB.MaxTries; i++ {
		log.Infof("Try connect to DB [%d]", i)

		db, err = gorm.Open("postgres", connectionString)
		if err == nil {
			break
		}

		log.Warnf("Couldn't connect to DB '%v'", err)
		log.Warnf("Sleep for %v", cfg.DB.Timeout)
		time.Sleep(cfg.DB.Timeout)
	}

	return db, err

}
