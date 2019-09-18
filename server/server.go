package main

import (
	"fmt"
	"path/filepath"
	. "github.com/jiangnanandi/go-a-project1/conf"
	"github.com/jiangnanandi/go-a-project1/httpman"
	"github.com/jiangnanandi/go-a-project1/server/router"

	log "github.com/cihub/seelog"
	"github.com/jiangnanandi/go-a-project1/utils"
)

func main() {
	conf, _ := filepath.Abs("./conf/server")
	utils.SetGlobalConfPath(conf)
	var appConf utils.ConfYaml
	var err error
	if appConf, err = utils.LoadConf(); err != nil {
		fmt.Println("Conf Error:", err.Error())
	}
	log.LoggerFromConfigAsFile("./conf/log.xml")
	defer log.Flush()
	httpconf := HttpConf{appConf.Http.Address, appConf.Http.Port, appConf.Http.ReadTimeout, appConf.Http.WriteTimeout}
	httpman.Run(router.SetupRouter, httpconf)
	log.Info("Server String...")
}
