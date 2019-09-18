package router

import (
	"github.com/jiangnanandi/go-a-project1/client/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(request *gin.Engine) *gin.Engine {
	request.GET("/list", controller.GetList)                               //静态页URL模板

	return request
}
