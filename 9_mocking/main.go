package main

import (
	"9_mocking/countdown"
	"9_mocking/sleeper"
	"os"
	"time"
)

func main() {
	sleeper := sleeper.NewConfigurableSleeper(1*time.Second, time.Sleep)

	countdown.Countdown(os.Stdout, sleeper)
}
