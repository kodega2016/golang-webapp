// Package config handle application wide configuration
package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLogger    *log.Logger
	InProduction  bool
	Sessoon       *scs.SessionManager
}
