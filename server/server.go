package main

import (
	"github.com/jiangnanandi/go-a-project1/httpman"
	"github.com/jiangnanandi/go-a-project1/server/router"

	log "github.com/cihub/seelog"
	"github.com/jiangnanandi/go-a-project1/utils"
)

func main() {
	utils.LoadConf()
	log.LoggerFromConfigAsFile("./conf/log.xml")
	defer log.Flush()
	httpman.Run(router.SetupRouter)
	log.Info("Server String...")
}
