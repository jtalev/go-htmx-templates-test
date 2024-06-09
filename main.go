package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("home.html"))
		
		data := Person{
			"Josh",
			"Talev",
			28,
		}
		tmpl.Execute(w, data)
	}

	http.HandleFunc("/", h1)

	http.ListenAndServe(":5000", nil)

}
