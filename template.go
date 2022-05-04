package main

import (
	"html/template"
	"log"
)

var (
	render *template.Template
)

func init() {
	log.Println("init template...")

	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		panic(err)
	}
	render = tmpl
}
