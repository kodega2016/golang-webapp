// Package render handles the rendering of the html templates
package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplateTest renders the html content using html template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("failed to execute the template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// variable to hold the template cache
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if the template is already available in the cache or not
	_, inMap := tc[t]
	if inMap {
		// template is already available in the template cache
		log.Println("template is already available in the cache")
	} else {
		log.Println("creating template")
		createTemplateCache(t)
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		"./templates/" + t,
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
