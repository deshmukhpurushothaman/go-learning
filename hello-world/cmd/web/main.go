package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/config"
	"github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/handlers"
	"github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/render"
)

const PORT = ":8000"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Could not create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
