package routers

import (
	"go-shop/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func findAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var t []model.Technology
	db.Find(&t)
	c.JSON(200, t)
}
