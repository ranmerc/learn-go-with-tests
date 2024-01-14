package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	endA := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	endB := time.Since(startB)

	if endA < endB {
		return a
	}

	return b
}
