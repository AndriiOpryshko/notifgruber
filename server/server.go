package server

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/AndriiOpryshko/notifgruber/notifications"
)

var notifProducer *notifications.NotificationProducer


func Run(addr string, np *notifications.NotificationProducer){
	log.WithFields(log.Fields{
		"addr": addr,
	}).Info("Server starts")

	notifProducer = np


	http.HandleFunc("/healthCheck", withTracing(HealthCheckHandler))
	http.HandleFunc("/notification", withTracing(NotificationHandler))

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop

		log.WithFields(log.Fields{
			"sig": sig,
			"wait_for_finish_sec": 2,
		}).Warning("Caught sig to finish")

		time.Sleep(2*time.Second)
		os.Exit(0)
	}()

	http.ListenAndServe(addr, nil)
}
