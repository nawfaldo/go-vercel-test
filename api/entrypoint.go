package api

import (
	"database/sql"
	"net/http"
	"vercer/service/user"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var (
	router *mux.Router
	db     *sql.DB
)

func init() {
	db, _ = sql.Open("postgres", "postgres://default:Ut4uNix0wdRk@ep-polished-sea-a1efivnq.ap-southeast-1.aws.neon.tech:5432/verceldb?sslmode=require")

	router = mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	user.RegisterRoutes(v1, db)

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
