package config

import "html/template"

// Template variable
var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}
