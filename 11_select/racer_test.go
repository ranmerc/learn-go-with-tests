package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers, returns one with fastest response", func(t *testing.T) {
		fastServer := makeDelayedServer(0 * time.Millisecond)
		slowServer := makeDelayedServer(20 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("didn't expect an error but got one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if neither respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(20 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)

		w.WriteHeader(http.StatusOK)
	}))
}
