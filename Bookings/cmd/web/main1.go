package main

import (
	"fmt"
	"github.com/suyashs52/golang/bookings/pkg/config"
	"github.com/suyashs52/golang/bookings/pkg/handlers"
	"github.com/suyashs52/golang/bookings/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main the the main application
// in WebBasic ....go run cmd/web/*.go
// go get -u github.com/go-chi/chi
// go mod tidy remove unused package
// go get github.com/justinas/nosurf for csrf
// go get github.com/alexedwards/scs/v2 for sessions
func main() {

	app.InProduction = false
	//true when in production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache1()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	//create a new repository
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
