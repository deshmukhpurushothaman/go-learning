package main

import (
	"fmt"
	"net/http"

	"github.com/deshmukhpurushothaman/go-learning/hello-world/pkg/handlers"
)

const PORT = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
