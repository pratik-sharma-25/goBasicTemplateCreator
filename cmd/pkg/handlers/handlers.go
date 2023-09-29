package handlers

import (
	"fmt"
	"net/http"

	"github.com/pratik25sharma/firstApi/cmd/pkg/config"
	"github.com/pratik25sharma/firstApi/cmd/pkg/models"
	"github.com/pratik25sharma/firstApi/cmd/pkg/renders"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// TemplateDate holds data set from the handlers to the template

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIp", remoteIp)
	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateDate{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello Pratik"

	fmt.Println("Hit the controller")

	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")

	stringMap["remoteIp"] = remoteIp

	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateDate{
		StringMap: stringMap,
	})
}
