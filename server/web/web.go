package web

import (
	"fmt"
	"go-shop/config"
	"go-shop/web/routers"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.Default())

	r.Use(func(c *gin.Context) {
		c.Set("db", w.db)
		c.Next()
	})

	routers.Register(r)
	return r.Run(fmt.Sprintf(":%d", w.cfg.Web.Port))
}
