package handlers

import (
	"net/http"

	"github.com/suyash/basicrest/pkg/config"
	"github.com/suyash/basicrest/pkg/models"
	"github.com/suyash/basicrest/pkg/render"
)

// Repo for repository used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handler
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	//perform the logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Again"

	//send the data to template
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Again"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
