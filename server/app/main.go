package main

import (
	"go-shop/config"
	"go-shop/db"
	"go-shop/web"

	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		cfg *config.Config
		err error
	)

	cfg, err = config.DefaultConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = mainFunc(cfg)
	if err != nil {
		log.Fatalln(err)
	}
}

func mainFunc(cfg *config.Config) error {
	db, err := db.NewPostgres(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	app, err := web.New(cfg, db)
	if err != nil {
		return err
	}
	return app.Serve()
}
