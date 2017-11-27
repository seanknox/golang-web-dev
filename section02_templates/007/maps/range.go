package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/friends.gotpl"))
}

func main() {
	friends := map[string]string{
		"Castro":    "Eve",
		"Sunnyside": "Dee",
		"Columbus":  "Alison",
		"Soma":      "Carcamo",
	}

	err := tpl.Execute(os.Stdout, friends)
	if err != nil {
		log.Fatalln(err)
	}

}
