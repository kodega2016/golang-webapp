package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kodega2016/booking-app/internal/pkg/config"
	"github.com/kodega2016/booking-app/internal/pkg/handlers"
	"github.com/kodega2016/booking-app/internal/pkg/models"
	"github.com/kodega2016/booking-app/internal/pkg/render"
)

var (
	portNumber = ":8080"
	app        config.AppConfig
	session    *scs.SessionManager
)

func main() {
	gob.Register(models.Reservation{})
	// setting up the session manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Application running on port %v\n", portNumber)

	srv := http.Server{
		Addr: portNumber, Handler: routes(&app),
	}

	srv.ListenAndServe()
}
