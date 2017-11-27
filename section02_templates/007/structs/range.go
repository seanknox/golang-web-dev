package main

import (
	"log"
	"os"
	"text/template"
)

type answer struct {
	Primary   string
	Secondary string
}

func main() {
	ans := answer{Primary: "42", Secondary: "monkey"}
	tpl, err := template.ParseFiles("templates/answer.gotpl")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, ans)
	if err != nil {
		log.Fatalln(err)
	}

}
