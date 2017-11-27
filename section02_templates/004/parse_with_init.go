package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	err := tpl.Execute(os.Stdout, nil)
	checkErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "foo.gotpl", nil)
	checkErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gotpl", nil)
	checkErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gotpl", nil)
	checkErr(err)

	// template.ExecuteTemplate returns an error without panicking, so
	// nothing will be returned to Stdout if the specified template file
	// is missing/bad.
	tpl.ExecuteTemplate(os.Stdout, "not_a_real_file.gotpl", nil)
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
