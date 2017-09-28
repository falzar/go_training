package main

import "html/template"
import "os"
import "log"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("go_web/templates/nested/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "Freddy is here!!")

	if err != nil {
		log.Fatalln(err)
	}
}
