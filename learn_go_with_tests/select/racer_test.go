package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatsusOK)
	}))

	fastServer := httptest.NewServer(http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatsusOK)
	}))

	slowURL := slowServer
	fastURL := fastServer

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
