package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/kodega2016/booking-app/internal/pkg/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Errorf("type is not allowed,%T", v)
	}
}
