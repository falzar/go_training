package main

import "html/template"
import "time"
import "os"
import "log"

var tpl *template.Template

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("go_web/templates/date_formatting/tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	// return t.Format("01-02-2006")
	// return t.Format(time.Kitchen)
	return t.Format(time.RFC1123)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())

	if err != nil {
		log.Fatalln(err)
	}
}
