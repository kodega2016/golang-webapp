// Package render handles the rendering of the html templates
package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders the html content using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("failed to execute the template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
