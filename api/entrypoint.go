package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
