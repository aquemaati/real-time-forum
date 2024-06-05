package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterData struct {
	Nickname  string `json:"nickname"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var data RegisterData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Unable to parse JSON data", http.StatusBadRequest)
			return
		}

		fmt.Println(data.Nickname)
		fmt.Println(data.Age)
		fmt.Println(data.Gender)
		fmt.Println(data.FirstName)
		fmt.Println(data.LastName)
		fmt.Println(data.Email)
		fmt.Println(data.Password)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
	})
}
