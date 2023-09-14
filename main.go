package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./index.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/read-my-joke", h1)
	log.Fatal(http.ListenAndServe(":8040", nil))
}
