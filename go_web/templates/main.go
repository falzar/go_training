package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

//Create a FuncMap to register functions.
// "uc" is what the func will be called in the template.
// "ft" is a function declared.

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type angel struct {
	// name        string -> Unexported field
	Name        string
	DestroyedBy string
}

//File level global variable. To be package level first letter shoul be uppercase
var gtpl *template.Template

func init() {
	// gtpl = template.Must(template.ParseFiles("go_web/templates/templates/tpl.gohtml", "go_web/templates/templates/wisdom.gohtml"))
	// gtpl = template.Must(template.ParseGlob("go_web/templates/templates/*.gohtml"))

	// IRC, template functions map must be defined by .Funcs before parsing the template.
	// Global template parse. parse all .gohtml files in folder go_web/templates/templates/
	// .New creates a new empty template
	gtpl = template.Must(template.New("").Funcs(fm).ParseGlob("go_web/templates/templates/*.gohtml")) //To insert funtions on templates
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	name := "John"

	tpl1 := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>SuperSite</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl1)

	// Template from file. Parse all listed files
	tpl2, err := template.ParseFiles("go_web/templates/templates/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl2.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Pass arguments to templates
	// Uses the {{.}} field on the template to locate data
	err = gtpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	// Pass arguments to templates and use un templates variables
	// Uses the {{.}} field on the template to locate data and then assigns to a variable within the template
	err = gtpl.ExecuteTemplate(os.Stdout, "wisdom.gohtml", `relese self focus; embrace other focus`)
	if err != nil {
		log.Fatalln(err)
	}

	// Pass a slice to templates
	s := []string{"Ramiel", "Azrael", "Lilith", "Sachiel", "Adam"}
	// Uses the {{range .}}<li>{{.}}</li>{{end}} field on the template to iterate over the slice
	err = gtpl.ExecuteTemplate(os.Stdout, "slice.gohtml", s)
	if err != nil {
		log.Fatalln(err)
	}

	// Pass a map to templates
	m := map[string]string{
		"Ramiel":  "Kendo",
		"Azrael":  "Asuka",
		"Lilith":  "Rei",
		"Sachiel": "Kaoru",
		"Adam":    "Eva",
	}

	// Uses the {{range .}}<li>{{.}}</li>{{end}} field on the template to iterate over the slice
	err = gtpl.ExecuteTemplate(os.Stdout, "map.gohtml", m)
	if err != nil {
		log.Fatalln(err)
	}

	// Pass a struct to templates
	r := angel{
		Name:        "Ramiel",
		DestroyedBy: "Eva 01",
	}

	// Uses the <li>{{.Name}} destroyed by {{.DestroyedBy}}</li> field on the template to show struct properties
	err = gtpl.ExecuteTemplate(os.Stdout, "struct.gohtml", r)
	if err != nil {
		log.Fatalln(err)
	}

	// Uses fm functions inside this template
	err = gtpl.ExecuteTemplate(os.Stdout, "functions.gohtml", r)
	if err != nil {
		log.Fatalln(err)
	}
}
