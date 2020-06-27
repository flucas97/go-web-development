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

	languages := []string{"Go", "Ruby", "React", "Typescript"}

	err := tpl.Execute(os.Stdout, languages)
	if err != nil {
		log.Fatalln(err)
	}
}
