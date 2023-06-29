package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		Addr:              ":8080",
		Handler:           mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		_ = server.Close()
	}()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello there!")
}
