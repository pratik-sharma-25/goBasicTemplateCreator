package main

import (
	"log"
	"net/http"

	"github.com/pratik25sharma/firstApi/cmd/pkg/config"
	"github.com/pratik25sharma/firstApi/cmd/pkg/handlers"
	"github.com/pratik25sharma/firstApi/cmd/pkg/renders"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template Cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	renders.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
