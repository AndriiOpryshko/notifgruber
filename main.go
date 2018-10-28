package main

import (
	"github.com/AndriiOpryshko/notifgruber/server"
	log "github.com/sirupsen/logrus"
	"fmt"
)


func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	InitConfig()
	addr := fmt.Sprintf("%s:%d", config.ApiConf.Addr, config.ApiConf.Port)
	server.Run(addr)
}
