package main

import (
	"8_dependency_injection/hello"
	"os"
)

func main() {
	hello.Greet(os.Stdout, "Elodie")
}
