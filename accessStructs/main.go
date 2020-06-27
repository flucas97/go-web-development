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

type cannabis struct {
	Name string
	Thc  int
}

func main() {
	flowers := []cannabis{
		{
			Name: "Gorila Haze",
			Thc:  22,
		},
		{
			Name: "Notherland",
			Thc:  15,
		},
	}

	err := tpl.Execute(os.Stdout, flowers)
	if err != nil {
		log.Fatalln(err)
	}
}
