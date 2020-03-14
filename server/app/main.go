package main

import (
	"go-shop/config"
	"go-shop/db"
	"go-shop/web"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "go-shop-server",
		Usage: "Shop server on Golang",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "profile",
				Value:   "dev",
				EnvVars: []string{"profile"},
				Usage:   "Current profile: prod or dev",
			},
		},
		Action: mainFunc,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainFunc(c *cli.Context) error {
	var (
		cfg *config.Config
		err error
	)

	profile := c.String("profile")
	log.Infof("Current profile is '%s'", profile)

	switch profile {
	case "prod":
		cfg, err = config.ProdConfig()
		break

	default:
		cfg, err = config.DevConfig()
		break
	}

	if err != nil {
		return err
	}

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
