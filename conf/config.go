package conf

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	PORT string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("conf/app.config.json", &conf)
	return conf
}

// var shared *_Configuration

// type _Configuration struct {
// 	Server struct {
// 		Port        int `json:"port`
// 		ReadTimeout time.Duration `json:"read_timeout"`
// 		WriteTimeout time.Duration`json:"write_timeout"`
// 	} `json:"server"`

// 	Log struct {
// 		Verbose bool `json:"verbose"`
// 	}`json:"log"`
// }