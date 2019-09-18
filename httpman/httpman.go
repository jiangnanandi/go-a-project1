package httpman

import (
	"github.com/jiangnanandi/go-a-project1/server/middleware"
	. "github.com/jiangnanandi/go-a-project1/conf"
	"github.com/jiangnanandi/go-a-project1/utils"
	log "github.com/cihub/seelog"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
	"syscall"
	"time"
)

func Run(inner func(r *gin.Engine) *gin.Engine, httpConf HttpConf) error {
	// http server
	appEngine := gin.New()
	appEngine.Use(middleware.Recovery())

	//初始化路由
	appRouterConfig := inner(appEngine)
	appServer := endless.NewServer(utils.MultiJoinString(httpConf.Address, ":", httpConf.Port), appRouterConfig)
	appServer.BeforeBegin = func(add string) {
		log.Info("Actual pid is ", syscall.Getpid())
	}
	//超时时间
	appServer.ReadTimeout = time.Duration(httpConf.ReadTimeout) * time.Second
	appServer.WriteTimeout = time.Duration(httpConf.WriteTimeout) * time.Second
	//监听http端口
	if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
