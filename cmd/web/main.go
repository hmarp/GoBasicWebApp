package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hmarp/basic-web-app/pkg/config"
	"github.com/hmarp/basic-web-app/pkg/handlers"
	"github.com/hmarp/basic-web-app/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app := config.AppConfig{
		UseCache:      false,
		TemplateCache: tc,
	}

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
