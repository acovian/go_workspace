package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

// method to save page.
func (p *Page) save() error {
	f := p.Title + ".txt"
	return ioutil.WriteFile(f, p.Body, 0600)
}

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view")]
	p, _ := load(title)
	t, _ := template.ParseFiles("test.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/view/", view)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/save/", save)
}
