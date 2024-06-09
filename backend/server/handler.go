package server

import (
	"fmt"

	"github.com/gorilla/mux"
)

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
