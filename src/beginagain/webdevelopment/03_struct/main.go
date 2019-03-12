package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type country struct {
	Country string
	Capital string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	mexico := country{
		Country: "Mexico",
		Capital: "Mexico City",
	}

	err := tpl.Execute(os.Stdout, mexico)
	if err != nil {
		log.Fatalln(err)
	}
}
