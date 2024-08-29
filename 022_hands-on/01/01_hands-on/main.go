package main

import (
	"io"
	"net/http"
)

func s(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "nothing")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark, bark")
}

func m(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Pasha")
}

func main() {
	http.HandleFunc("/", s)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8085", nil)
}
