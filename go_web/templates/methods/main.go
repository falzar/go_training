package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type person struct {
	Name string
	Age  int
}

func (p person) AgeDouble() int {
	return p.Age * 2
}

// func (p person) nameToUpercase() string { Method local to this file
func (p person) NameToUpercase() string {
	return strings.ToUpper(p.Name)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("go_web/templates/methods/tpl.gohtml"))
}

func main() {

	jim := person{
		Name: "Jim",
		Age:  15,
	}

	err := tpl.Execute(os.Stdout, jim)

	if err != nil {
		log.Fatalln(err)
	}
}
