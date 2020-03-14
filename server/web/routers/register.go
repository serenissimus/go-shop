package routers

import (
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	e.GET("/technologies", findAll)
}
