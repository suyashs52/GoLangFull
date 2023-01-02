package handlers

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
	"github.com/suyashs52/golang/bookings/internal/config"
	"github.com/suyashs52/golang/bookings/internal/models"
	"github.com/suyashs52/golang/bookings/internal/render"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

func geRoutes() http.Handler {
	//what am i going to put in the session
	gob.Register(models.Reservation{})

	app.InProduction = false
	//true when in production

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErroLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCache1()

	if err != nil {
		log.Fatal("cannot create template cache", err)

	}

	app.TemplateCache = tc
	app.UseCache = true

	//create a new repository
	repo := NewRepo(&app)
	NewHandler(repo)

	render.NewTemplates(&app)
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	//mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)
	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/contact", Repo.Contact)
	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

// NoSurf add CSRF protestion to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

var functions = template.FuncMap{}
var pathToTemplate = "./../../templates"

func CreateTestTemplateCache1() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files names *page.tmpl  from ./Template

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplate))
	if err != nil {
		return myCache, err
	}

	// range through al files ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*layout.tmpl", pathToTemplate))

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*layout.tmpl", pathToTemplate))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}
