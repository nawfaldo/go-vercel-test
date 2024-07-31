package api

import (
	"net/http"
	"vercer/service/user"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	user.RegisterRoutes(v1)

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
