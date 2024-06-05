package server

import (
	"encoding/json"
	"net/http"

	"github.com/aquemaati/real-time-forum/handlers"
	"github.com/aquemaati/real-time-forum/middleware"
)

type Message struct {
	Message string `json:"message"`
}

// Routing configure les routes de l'application et retourne un http.Handler.
func Routing() http.Handler {
	mux := http.NewServeMux()

	// Utilisation du middleware CORS
	handler := middleware.CORSMiddleware(mux)

	// Servir les fichiers statiques du répertoire frontend/src
	mux.Handle("/", http.FileServer(http.Dir("./frontend/src")))

	// Endpoint API qui retourne un message JSON
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Message{Message: "Hello, world!"})
	})

	// Route vers le gestionnaire Home
	mux.HandleFunc("/home", handlers.Home)

	// Route vers le gestionnaire Register
	mux.Handle("/api/register", handlers.Register())

	return handler
}
