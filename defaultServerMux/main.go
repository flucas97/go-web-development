package main

import (
	"io"
	"log"
	"net/http"
)

func way(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "there is still a way")
}

func noway(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "no way to do that dude")
}

func main() {
	http.HandleFunc("/way", way)
	http.HandleFunc("/noway", noway)

	// passing nil to use Handle and HandleFunc from http package
	log.Fatal(http.ListenAndServe(":80", nil)) // default server mux
}
