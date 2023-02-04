package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// Since function name starts in capital letter,
// it can be called from another package
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}
