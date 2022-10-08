package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/vietthangc1/booking/internal/config"
	"github.com/vietthangc1/booking/internal/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(data *models.TemplateData, req *http.Request) *models.TemplateData {
	data.CSRFToken = nosurf.Token(req)
	return data
}

func RenderTemplate(w http.ResponseWriter, req *http.Request, tmpl string, data *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	parsedTemplate, check := templateCache[tmpl]
	if !check {
		log.Fatal("Error")
	}

	buffer := new(bytes.Buffer)
	data = AddDefaultData(data, req)
	err := parsedTemplate.Execute(buffer, data)
	if err != nil {
		log.Println(err)
	}

	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	return
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	var myCache = make(map[string]*template.Template)

	// lấy template html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		parsedTemplate, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// lấy layout html
		layouts, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			parsedTemplate, err = parsedTemplate.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = parsedTemplate
	}
	return myCache, nil
}

// Cách 1 Template cache

// var templateCache = make(map[string]*template.Template)
// func RenderTemplate(w http.ResponseWriter, t string)  {
// 	var tmpl *template.Template

// 	_, inMap := templateCache[t]

// 	if (!inMap) {
// 		// chưa có, cần thên
// 		log.Println("Add new template")
// 		parsedTemplate, err := template.ParseFiles("./templates/" + t, "./templates/base.layout.html")
// 		if (err != nil) {
// 			fmt.Println("Error in add new template", err)
// 		}
// 		templateCache[t] = parsedTemplate
// 	} else {
// 		log.Println("Using cached")
// 	}
// 	tmpl = templateCache[t]
// 	err := tmpl.Execute(w, nil)
// 	if (err != nil) {
// 		fmt.Println("Error in execute",err)
// 	}
// 	log.Println(templateCache)
// 	return
// }
