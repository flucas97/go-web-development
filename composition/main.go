package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type person struct {
	Name     string
	Lastname string
}

type secretAgent struct {
	person
	LicenseToKill bool
}

func main() {
	james := secretAgent{
		person{
			Name:     "James",
			Lastname: "Bound",
		},
		true,
	}

	err := tpl.Execute(os.Stdout, james)
	if err != nil {
		log.Fatalln(err)
	}
}
