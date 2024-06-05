package handlers

import (
	"fmt"
	"net/http"
)

func Register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form data", http.StatusBadRequest)
			return
		}

		// Retrieve form values
		nickname := r.FormValue("nickname")
		age := r.FormValue("age")
		gender := r.FormValue("gender")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		email := r.FormValue("email")
		password := r.FormValue("password")

		fmt.Println(nickname)
		fmt.Println(age)
		fmt.Println(gender)
		fmt.Println(firstName)
		fmt.Println(lastName)
		fmt.Println(email)
		fmt.Println(password)
	})
}
