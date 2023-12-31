package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/pratik25sharma/firstApi/cmd/pkg/config"
	"github.com/pratik25sharma/firstApi/cmd/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateDate) *models.TemplateDate {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateDate) {
	var tc map[string]*template.Template
	// get the template cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template")
	}

	buf := new(bytes.Buffer)

	td = addDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	// render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files with tmpl file
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range throug the pages
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
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

	return myCache, err
}
