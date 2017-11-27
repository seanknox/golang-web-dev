package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("templates/*")
	checkErr(err)

	err = tpl.Execute(os.Stdout, nil)
	checkErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gotpl", nil)
	checkErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gotpl", nil)
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
