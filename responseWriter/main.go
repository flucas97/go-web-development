package main

import (
	"fmt"
	"net/http"
)

type workflow int

func (workflow workflow) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Feijo-Key", "this is from flucas97")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var w workflow
	http.ListenAndServe(":8080", w)
}
