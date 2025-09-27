// Package handlers handles the web application request
package handlers

import (
	"net/http"

	"github.com/kodega2016/booking-app/pkg/config"
	"github.com/kodega2016/booking-app/pkg/models"
	"github.com/kodega2016/booking-app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(config *config.AppConfig) *Repository {
	return &Repository{
		App: config,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (*Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (*Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
