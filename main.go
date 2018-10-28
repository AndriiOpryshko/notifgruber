package main

import (
	"github.com/AndriiOpryshko/notifgruber/server"
)


func main() {
	addr := ":8080"
	server.Run(addr)
}
