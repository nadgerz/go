package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Millisecond)
		serverB := makeDelayedServer(12 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

	t.Run("compares two servers and returns the faster one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an eeor but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
