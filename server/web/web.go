package web

import (
	"fmt"
	"go-shop/config"
	"go-shop/web/router"

	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Web struct {
	db  *gorm.DB
	cfg *config.Config
}

func New(cfg *config.Config, db *gorm.DB) (*Web, error) {
	web := &Web{
		db:  db,
		cfg: cfg,
	}
	return web, nil
}

func (w *Web) Serve() error {
	r := gin.Default()

	if w.cfg.Web.Cors {
		r.Use(cors.Default())
	}

	r.Use(func(c *gin.Context) {
		c.Set("db", w.db)
		c.Next()
	})

	r.Use(static.Serve("/", static.LocalFile("/webapp", false)))

	log.Info("Register routes")
	router.Register(r)

	log.Infof("Launch server on port %d", w.cfg.Web.Port)
	return r.Run(fmt.Sprintf(":%d", w.cfg.Web.Port))
}
