package web

import (
	"go-shop/config"
	"go-shop/web/routers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Web struct {
	DB *gorm.DB
}

func New(cfg *config.Config, db *gorm.DB) (*Web, error) {
	web := &Web{
		DB: db,
	}
	return web, nil
}

func (w *Web) Serve() error {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", w.DB)
		c.Next()
	})
	routers.Register(r)
	return r.Run()
}
