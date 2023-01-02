package render

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/suyashs52/golang/bookings/internal/config"
	"github.com/suyashs52/golang/bookings/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// Render template  renders template using html/template
func RenderTemplateTest(w http.ResponseWriter, temp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+temp, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest1(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	//
	//check to see if we already have template in our chache

	_, inMap := tc[t]
	if !inMap {
		//need to create the template
		err = createTemplateCache(t)

		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cache template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}
	//parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err

	}

	tc[t] = tmpl
	return nil
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, temp string, td *models.TemplateData) error {
	//create template cached
	//tc, err := CreateTemplateCache1()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache1()
	}

	//get requested templae from cache
	t, ok := tc[temp]

	if !ok {

		log.Println("could not create template cache")
		//	return errors.New("can't create template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

var pathToTemplate = "./templates"
var functions = template.FuncMap{}

func CreateTemplateCache1() (map[string]*template.Template, error) {
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
