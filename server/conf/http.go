package conf

type HttpConf struct {
	Address      string
	Port         string
	ReadTimeout  int
	WriteTimeout int
}