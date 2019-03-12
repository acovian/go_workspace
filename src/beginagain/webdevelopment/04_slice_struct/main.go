package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type trilobite struct {
	Name   string
	Period string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	asaphicus := trilobite{
		Name:   "Asaphicus",
		Period: "Cambrian",
	}
	acidaspis := trilobite{
		Name:   "Acidaspis",
		Period: "Ordovician",
	}
	phacops := trilobite{
		Name:   "Phacops",
		Period: "Devonian",
	}
	trilobites := []trilobite{asaphicus, acidaspis, phacops}
	err := tpl.Execute(os.Stdout, trilobites)
	if err != nil {
		log.Fatalln(err)
	}
}
