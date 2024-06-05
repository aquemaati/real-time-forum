package server

import (
	"net/http"

	"github.com/aquemaati/real-time-forum/handlers"
	"github.com/aquemaati/real-time-forum/middleware"
)

func ChainMiddle() {

}
func Routing() http.Handler {
	mux := http.NewServeMux()
	handler := middleware.CORSMiddleware(mux)

	mux.Handle("/register", handlers.Register())
	return handler
}
