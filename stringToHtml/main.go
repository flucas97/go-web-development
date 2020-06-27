package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Lucas"

	tpl := `
	<!DOCTYPE html>
	<head>
	<meta charset="UTF-8">
	<title> Hello world! </title>
	</head>
	<body>
	<h1> Hello, ` + name + `</h1>
	</body>
	</html>
	`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error")
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
