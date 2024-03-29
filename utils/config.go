package utils

import (
	"fmt"
	"path/filepath"
	"github.com/spf13/viper"
	log "github.com/cihub/seelog"
)

type ConfYaml struct {
	Http SectionHttp `yaml:"http"`
}

type SectionHttp struct {
	Enabled      bool   `yaml:"enabled"`
	Address      string `yaml:"address"`
	Port         string `yaml:"port"`
	ReadTimeout  int    `yaml:"readtimeout"`
	WriteTimeout int    `yaml:"writetimeout"`
}

/*资源预加载*/
func LoadConf(basePath string){
	confPath := basePath + "/conf"
	absPath, _ := filepath.Abs(confPath)
	fmt.Println(absPath + `/config.json`)
	viper.SetConfigFile(absPath + `/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}
}
