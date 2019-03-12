package main

import (
	"html/template"
	"log"
	"os"
)

type car struct {
	Model, Year string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	cars := []car{
		car{
			Model: "Jaguar f-type",
			Year:  "2017",
		},
		car{
			Model: "BMW Convertible",
			Year:  "2016",
		},
	}

	err := tpl.Execute(os.Stdout, cars)
	if err != nil {
		log.Fatalln(err)
	}
}
