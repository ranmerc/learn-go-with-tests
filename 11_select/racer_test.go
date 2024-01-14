package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastServer := makeDelayedServer(0 * time.Millisecond)

	slowServer := makeDelayedServer(20 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL

	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)

		w.WriteHeader(http.StatusOK)
	}))
}
