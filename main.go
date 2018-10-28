package main

import (
	"github.com/AndriiOpryshko/notifgruber/server"
	"github.com/AndriiOpryshko/notifgruber/notifications"
	log "github.com/sirupsen/logrus"
	"fmt"
)


func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	InitConfig()
	notsProducer := notifications.InitProducer(config.KafkaConf.Addr, config.KafkaConf.NotifTopic)
	addr := fmt.Sprintf("%s:%d", config.ApiConf.Addr, config.ApiConf.Port)

	go notsProducer.Run()
	server.Run(addr, notsProducer)
}
