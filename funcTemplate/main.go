package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"sayHi": sayHi,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

func sayHi() string {
	return "Hello!"
}
