package handlers

import (
	"fmt"
	"html"
	"net/http"
)

// Function where we can see all posts
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
