package utils

import (
	"fmt"

	"github.com/spf13/viper"
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

var configPath string

func SetGlobalConfPath(path string) {
	configPath = path
}

func GetGlobalConfPath() string {
	return configPath
}

func LoadConf() (ConfYaml, error) {
	var conf ConfYaml
	//设定文件格式
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	// 打开配置
	if err := viper.ReadInConfig(); err != nil {
		return conf, err
	} else {
		fmt.Println("Using config.go file:", viper.ConfigFileUsed())
	}

	// Web
	conf.Http.Address = viper.GetString("Http.address")
	conf.Http.Port = viper.GetString("Http.port")
	conf.Http.ReadTimeout = viper.GetInt("Http.readtimeout")
	conf.Http.WriteTimeout = viper.GetInt("Http.writetimeout")

	return conf, nil
}
