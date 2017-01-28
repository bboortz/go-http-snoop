package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	var ipport string = ":" + port
	router := NewRouter()

	log.Info("listening on: ", ipport)
	log.Fatal(http.ListenAndServe(ipport, router))
}
