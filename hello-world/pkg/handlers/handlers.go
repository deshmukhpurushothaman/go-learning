package handlers

import (
	"net/http"

	"github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/config"
	"github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/models"
	"github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/render"
	// "github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/render"
	// "github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets new repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
