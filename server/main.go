package main

import "net/http"

func main() {
	server := http.Server{}
	_ = server.ListenAndServe()

	defer func() {
		_ = server.Close()
	}()
}
