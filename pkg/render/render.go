// Package render handles the rendering of the html templates
package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/kodega2016/booking-app/pkg/config"
)

var app *config.AppConfig

// NewTemplates will set the app config
func NewTemplates(newConfig *config.AppConfig) {
	app = newConfig
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc := app.TemplateCache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("failed to get the template:")
	}

	buff := new(bytes.Buffer)
	err := t.Execute(buff, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all pages from the templates directory
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// loop through the pages and create template
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return nil, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
