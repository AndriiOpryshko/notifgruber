package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	log "github.com/sirupsen/logrus"
)

// For testing healthcheck and notification requests
type responseExpected struct {
	respTestName string
	code         int
	body         string
}

// expected response for healthCheck
var healthCheckTestResp = responseExpected{
	respTestName: "test healthCheck",
	code:         200,
	body:         `{"success":true,"message":""}`,
}

// Test healthCheck handler
func TestHealthCheckHandler(t *testing.T) {
	log.WithFields(log.Fields{
		"testName": healthCheckTestResp.respTestName,
	}).Info("Testing healthCheck")

	req, err := http.NewRequest("GET", "/healthCheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != healthCheckTestResp.code {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != healthCheckTestResp.body{
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), healthCheckTestResp.body)
	}
}

// expected responce for notifications
var notsTestResps = []responseExpected{
	{
		respTestName: "test notification bad structure",
		code:         400,
		body:         `{"success":false,"message":"Wrong structure of body"}`,
	},
	{
		respTestName: "test success notification post",
		code:         200,
		body:         `{"success":true,"message":""}`,
	},
}

// posted notifications for tests
var postedNots = []string{
	"test",
	`[{"id":"777","provider":"Cyren","service":"untiphishing"}]`,
}

// Test notification handler
func TestNotificationHandler(t *testing.T) {
	for i, nots := range postedNots {

		log.WithFields(log.Fields{
			"testName": notsTestResps[i].respTestName,
		}).Info("Testing notification")

		req, err := http.NewRequest("POST", "/notification", strings.NewReader(nots))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NotificationHandler)

		handler.ServeHTTP(rr, req)

		expectedStatus := notsTestResps[i].code
		if status := rr.Code; status != expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := notsTestResps[i].body
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}

	}
}
