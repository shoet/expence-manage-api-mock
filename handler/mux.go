package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewMux() http.Handler {
	router := chi.NewRouter()

	le := ListExpenses{}
	router.Get("/expenses", le.ServeHTTP)

	return router

}
