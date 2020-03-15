package main

import (
	"html/template"
	"net/http"
)

type IndexHtml struct {
	ClaimcheckReachable string
	Aantal int
}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.html")
	if (err != nil) {
		panic(err)
	}
	err = t.Execute(w, IndexHtml{ ClaimcheckReachable: "Ja", Aantal: 42 })
	if (err != nil) {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
