package httpman

import (
	"github.com/jiangnanandi/go-a-project1/server/middleware"
	log "github.com/cihub/seelog"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
	"syscall"
	"time"
	"github.com/spf13/viper"
)

type server struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
}

func Run(inner func(r *gin.Engine) *gin.Engine) error {
	s := &server{
		Address:      viper.GetString(`server.address`),
		ReadTimeout:  viper.GetInt64(`server.readtimeout`),
		WriteTimeout: viper.GetInt64(`server.writetimeout`),
	}
	// http server
	appEngine := gin.New()
	appEngine.Use(middleware.Recovery())

	//初始化路由
	appRouterConfig := inner(appEngine)
	appServer := endless.NewServer(s.Address, appRouterConfig)
	appServer.BeforeBegin = func(add string) {
		log.Info("Actual pid is ", syscall.Getpid())
	}
	//超时时间
	appServer.ReadTimeout = time.Duration(s.ReadTimeout) * time.Second
	appServer.WriteTimeout = time.Duration(s.WriteTimeout) * time.Second
	//监听http端口
	if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
