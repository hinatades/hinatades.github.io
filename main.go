package main

import (
	"html/template"
	"log"
	"net/http"
)

func handleHTML(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))

	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handleHTML)
	http.ListenAndServe(":8000", nil)
}
