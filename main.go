package main

import (
	"github.com/AndriiOpryshko/notifgruber/server"
	log "github.com/sirupsen/logrus"
)


func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	addr := ":8080"
	server.Run(addr)
}
