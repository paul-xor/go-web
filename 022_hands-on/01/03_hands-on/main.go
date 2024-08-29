package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func s(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "nothing")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark, bark")
}

func pasha(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", "Pasha")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", s)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", pasha)

	http.ListenAndServe(":8085", nil)
}
