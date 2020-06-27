package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	colors := map[string]string{
		"a": "red",
		"b": "yellow",
	}

	err := tpl.Execute(os.Stdout, colors)
	if err != nil {
		log.Fatalln(err)
	}
}
