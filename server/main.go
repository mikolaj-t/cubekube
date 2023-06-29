package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const Timeout = 10 * time.Second

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		ReadHeaderTimeout: Timeout,
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

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello there!")
}
