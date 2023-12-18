package main

import (
	"8_dependency_injection/hello"
	"log"
	"net/http"
	"os"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	hello.Greet(w, "Hey There!")
}

func main() {
	hello.Greet(os.Stdout, "Elodie")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
