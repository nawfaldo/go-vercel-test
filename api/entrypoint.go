package api

import (
	"net/http"
	"vercer/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/hello", hello).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://go-react-api-web.vercel.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.Handle("/", handler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {

	hello := map[string]string{
		"msg": "hello",
	}

	utils.WriteJSON(w, http.StatusAccepted, hello)
}
