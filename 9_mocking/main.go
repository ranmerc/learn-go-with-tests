package main

import (
	"9_mocking/countdown"
	"9_mocking/sleeper"
	"os"
)

func main() {
	sleeper := &sleeper.DefaultSleeper{}

	countdown.Countdown(os.Stdout, sleeper)
}
