package main

import (
	"github.com/jiangnanandi/go-a-project1/client/router"
	"github.com/jiangnanandi/go-a-project1/httpman"

	log "github.com/cihub/seelog"
	"github.com/jiangnanandi/go-a-project1/utils"
)

func main() {
	utils.LoadConf("/app/client")
	log.LoggerFromConfigAsFile("./conf/log.xml")
	defer log.Flush()
	httpman.Run(router.SetupRouter)
	log.Info("Server String...")
}
