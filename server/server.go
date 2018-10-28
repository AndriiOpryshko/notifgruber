package server

import (
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(addr string){
	log.Printf("Server starts on port: %s\n", addr)
	http.HandleFunc("/healthCheck", withTracing(HealthCheckHandler))
	http.HandleFunc("/notification", withTracing(NotificationHandler))

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("Caught sig: %+v\n", sig)
		log.Println("Wait for 2 second to finish processing\n")
		time.Sleep(2*time.Second)
		os.Exit(0)
	}()

	http.ListenAndServe(addr, nil)
}
