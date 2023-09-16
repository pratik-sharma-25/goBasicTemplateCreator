package handlers

import (
	"net/http"

	"github.com/pratik25sharma/firstApi/cmd/pkg/renders"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.tmpl")
}
