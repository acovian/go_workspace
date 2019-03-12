package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// parse file
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute file
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Release self-focus; embrace other-focus.")
	if err != nil {
		log.Fatalln(err)
	}
}
