package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/suyash/basicrest/pkg/config"
	"github.com/suyash/basicrest/pkg/models"
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

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, temp string, td *models.TemplateData) {
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
		log.Fatal("could not create template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache1() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files names *page.tmpl  from ./Template

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through al files ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}
