package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", home{})
	mux.HandleFunc("/posts", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Visit http://localhost:8080/posts to get started"))
	})
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

type home struct{}

func (h home) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello World! Welcome to the \"Just Enough Go\" blog series!!!"))
}
