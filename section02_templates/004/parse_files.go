package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("templates/foo.gotpl")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
