package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vietthangc1/booking/pkg/config"
	"github.com/vietthangc1/booking/pkg/handlers"
	"github.com/vietthangc1/booking/pkg/render"
)

const port = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false
	
	session = scs.New()
	session.Lifetime = 24*time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = templateCache
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Running at http://localhost%s", port))

	service := &http.Server{
		Addr: "localhost"+port,
		Handler: routes(&app),
	}

	err = service.ListenAndServe()
	log.Fatal(err)
}
