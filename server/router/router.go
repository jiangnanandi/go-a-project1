package router

import (
	"github.com/jiangnanandi/go-a-project1/server/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(request *gin.Engine) *gin.Engine {
	request.GET("/user", controller.GetUser)                               //静态页URL模板

	return request
}
