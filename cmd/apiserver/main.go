package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Fagan04/http-rest-api/app/config"
	"github.com/Fagan04/http-rest-api/app/server"
	log "github.com/sirupsen/logrus"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".toml", "path to config file")
}

func main() {
	flag.Parse()

	newConfig := config.NewConfig()
	_, err := toml.DecodeFile(configPath, &newConfig)
	if err != nil {
		log.Error("Could not read config file: ", err)
		return
	}

	s := server.NewAPIServer(newConfig)
	err = s.Start()
	if err != nil {
		log.Error("Could not start server: ", err)
		return
	}
}
