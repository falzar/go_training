package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqrRoot,
}

func double(i int) int {
	return 2 * i
}

func square(i int) int {
	return i * i
}

func sqrRoot(i int) float64 {
	return math.Sqrt(float64(i))
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("go_web/templates/pipeline/tpl.gohtml"))
}

func main() {

	//Pipeline helps to chain funtions on templates. For this template:
	// <h1>T{{. | fdbl}}</h1>
	// <h1>T{{. | fdbl | fsq}}</h1>
	// <h1>T{{. | fdbl | fsq | fsqrt}}</h1>
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 5)

	if err != nil {
		log.Fatalln(err)
	}

}
