package main

import (
	"log"
	"net/http"

	"github.com/willian2s/learn-go-with-test/di"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "word")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
