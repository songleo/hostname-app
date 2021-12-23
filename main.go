package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	socket := fmt.Sprintf("0.0.0.0:%v", port)
	log.Println(fmt.Sprintf("hostname server running on %v ...", socket))
	log.Fatal(http.ListenAndServe(socket, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintln(w, fmt.Sprintf("hostname: %v", host))

	version := os.Getenv("VERSION")
	if version == "" {
		version = "latest"
	}

	fmt.Fprintln(w, fmt.Sprintf("app version: %v", version))
	fmt.Fprintln(w, fmt.Sprintf("\n"))
	log.Println(fmt.Sprintf("method=%s url=%s protocol=%s status_code=%d",
		r.Method,
		r.URL,
		r.Proto,
		200))
}
