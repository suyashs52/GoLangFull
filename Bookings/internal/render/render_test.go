package render

import (
	"github.com/suyashs52/golang/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash 123 not found in session data")
	}

}

func TestRenderTemplate(t *testing.T) {
	pathToTemplate = "./../../templates"
	tc, err := CreateTemplateCache1()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})

	if err != nil {
		t.Error("error writing template to browser")
	}
	//err = RenderTemplate(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	//
	//if err == nil {
	//	t.Error("error not writing to template to browser")
	//}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))

	r = r.WithContext(ctx)

	return r, err
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache1(t *testing.T) {
	//pathToTemplate:="./../../templates"

	_, err := CreateTemplateCache1()

	if err != nil {
		t.Error(err)
	}
}
