package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type toad struct {
	Name  string
	Color string
}

func init() {
	tpl = template.Must(template.ParseGlob("toadsworth.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	toadsworth := toad{
		Name:  "Toadsworth",
		Color: "Brown",
	}
	toadette := toad{
		Name:  "Toadette",
		Color: "Pink",
	}
	toads := []toad{toadsworth, toadette}
	tpl.ExecuteTemplate(w, "toadsworth.html", toads)
}
