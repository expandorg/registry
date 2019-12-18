package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func withHandlers(r http.Handler) http.Handler {
	return loggingHandler(r)
}

func loggingHandler(r http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, r)
}
