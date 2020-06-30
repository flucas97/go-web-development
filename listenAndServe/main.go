package main

import (
	"fmt"
	"net/http"
)

type workflow int

func (workflow workflow) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var w workflow
	http.ListenAndServe(":8080", w)
}
