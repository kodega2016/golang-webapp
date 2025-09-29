// Package models holds the data models for the template
package models

import (
	"github.com/kodega2016/booking-app/internal/pkg/forms"
)

// TemplateData holds the data set sent from the handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
