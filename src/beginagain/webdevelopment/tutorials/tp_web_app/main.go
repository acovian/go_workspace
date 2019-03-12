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

// saves data into text files. Method.
func (p *Page) save() error {
	f := p.Title + ".txt"
	return ioutil.WriteFile(f, p.Body, 0600)
}

// loads text files. Outputs page, error.
func load(title string) (*Page, error) {
	f := title + ".txt"
	body, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// func to put all into server and display it
func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/test/"):]
	p, _ := load(title)
	t, _ := template.ParseFiles("test.html")
	t.Execute(w, p)
}

func edit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, _ := load(title)
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

// takes value from form which is inputted, replaces with
// title and body, and resaving back into text file, redirects
// us back to original test page.
func save(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/test/"+title, http.StatusFound)
}

func main() {
	p := &Page{Title: "Test", Body: []byte("Welcome to the Test Page.")}
	p.save()
	http.HandleFunc("/test/", view)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/save/", save)
	http.ListenAndServe(":3000", nil)
}
