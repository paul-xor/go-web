package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func s(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "nothing")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bark, bark")
}

func pasha(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", "Pasha")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(s))
	http.HandleFunc("/dog/", http.HandlerFunc(d))
	http.HandleFunc("/me/", http.HandlerFunc(pasha))

	http.ListenAndServe(":8085", nil)
}
