package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-backend/backend/jsonfiles"
	"real-time-backend/backend/modals"

	"github.com/gorilla/mux"
)

func categorieHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := jsonfiles.GetCategoriesTable()
	if err != nil {
		http.Error(w, "Error fetching categories", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := jsonfiles.GetPostTable()
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	users, err := jsonfiles.GetUserTable()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := modals.Response{Message: "Welcome"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LoadServer starts the server
func LoadServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/api/users", userHandler).Methods("GET")
	router.HandleFunc("/api/posts", postHandler).Methods("GET")
	router.HandleFunc("/api/categories", categorieHandler).Methods("GET")

	serverConfig := ServerParameters(router, 10)
	fmt.Println("Server started at 127.0.0.1:8080")
	if err := serverConfig.ListenAndServe(); err != nil {
		fmt.Println("Error to serve:", err)
	}
}
