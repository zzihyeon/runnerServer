package router

import (
	"runnerserver/config"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Initialize() {
	r = gin.Default()
	registerAPI()
}

func Run() {
	r.Run(":" + config.Development.APIPort)
}

func registerAPI() {
	r.GET("/ping", Ping)
}
