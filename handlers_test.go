package main

import (
	"github.com/chaselengel/checkin/mailer"
	"github.com/chaselengel/checkin/nodestore"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Init global database
func TestMain(m *testing.M) {
	var err error
	ns, err = nodestore.Open()
	if err != nil {
		panic(err)
	}
	mail, err = mailer.Open()
	if err != nil {
		panic(err)
	}
	// Run tests
	m.Run()
}

// Replace /script/checkin/{scriptName}
// with /script/checkin/test for all routes
// If endpoint doesn't have any replacments return unchanged endpoint
func replace(endpoint string) string {
	apiVars := map[string]string{
		"{scriptId}":   "1",
		"{scriptName}": "test",
		"{mailId}":     "test@example.com",
	}
	for key, val := range apiVars {
		newStr := strings.Replace(endpoint, key, val, -1)
		if newStr != endpoint {
			return newStr
		}
	}
	return endpoint
}

// Test that all routes return status 200
func TestStatusOK(t *testing.T) {
	for _, route := range routes {
		endpoint := replace(route.Path)
		req, err := http.NewRequest(route.Method, endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(route.HandlerFunc)
		handler.ServeHTTP(recorder, req)
		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("Handler return incorrect status:\nexpected %v\nactual%v", http.StatusOK, status)
		}
	}
}

// Test that all routes return application/json in Header
func TestHeaderIsJSON(t *testing.T) {
	for _, route := range routes {
		endpoint := replace(route.Path)
		req, err := http.NewRequest(route.Method, endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(route.HandlerFunc)
		handler.ServeHTTP(recorder, req)
		if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
			t.Errorf("Content-Type for %v was %v not application/json", endpoint, contentType)
		}
	}
}
