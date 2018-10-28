package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/AndriiOpryshko/notifgruber/notifications"
	log "github.com/sirupsen/logrus"
)

const (
	EMPTYBODYREQUEST = "Requested body is empty"
	WRONGBODYSTRUCT  = "Wrong structure of body"
)

type response struct {
	IsSuccess bool   `json:"success"'`
	Msg       string `json:"message"`
}

// Health check handler
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp := response{
		IsSuccess: true,
	}

	wrightResponse(w, http.StatusOK, resp)
}

// Notification handler
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	var resp response

	if err != nil {
		resp.IsSuccess = false
		resp.Msg = EMPTYBODYREQUEST
		log.Printf("%s. %s", resp.Msg, err)

		log.WithFields(log.Fields{
			"msg": resp.Msg,
			"err": err,
		}).Error("Error while gets body")

		wrightResponse(w, http.StatusBadRequest, resp)
		return
	}

	var nots []notifications.Notification

	err = json.Unmarshal(bodyBytes, &nots)

	log.WithFields(log.Fields{
		"notifications": nots,
	}).Debug("Posted notifications")

	if err != nil {
		resp.IsSuccess = false
		resp.Msg = WRONGBODYSTRUCT
		log.WithFields(log.Fields{
			"msg": resp.Msg,
			"err": err,
		}).Error("Error while gets body")
		wrightResponse(w, http.StatusBadRequest, resp)
		return
	}

	resp.IsSuccess = true
	wrightResponse(w, http.StatusOK, resp)
}

// Wrighte response code and body
func wrightResponse(w http.ResponseWriter, code int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}
