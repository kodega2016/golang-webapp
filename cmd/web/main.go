package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kodega2016/booking-app/pkg/config"
	"github.com/kodega2016/booking-app/pkg/handlers"
	"github.com/kodega2016/booking-app/pkg/render"
)

var portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application running on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
