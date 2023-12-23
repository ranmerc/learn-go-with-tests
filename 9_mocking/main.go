package main

import (
	"9_mocking/countdown"
	"os"
)

func main() {
	countdown.Countdown(os.Stdout)
}
