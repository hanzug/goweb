package routes

import (
	"github.com/gin-gonic/gin"
	"goweb/controller"
	"goweb/logger"
	"net/http"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signup", controller.SignUpHandler)
	r.POST("login", controller.LoginHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return r
}
