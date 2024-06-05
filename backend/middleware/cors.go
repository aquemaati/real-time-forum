package middleware

import "net/http"

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// // Add CORS headers
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// // Handle preflight request
		// if r.Method == http.MethodOptions {
		// 	w.WriteHeader(http.StatusNoContent)
		// 	return
		// }

		// Pass to the next handler
		next.ServeHTTP(w, r)
	})
}
