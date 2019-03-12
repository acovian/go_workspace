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
	capitals := map[string]string{
		"Iran":   "Tehran",
		"Turkey": "Ankara",
		"Egypt":  "Cairo",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", capitals)
	if err != nil {
		log.Fatalln(err)
	}
}
