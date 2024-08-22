package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// ParseFiles has variadic param and can parse mutiple files
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
