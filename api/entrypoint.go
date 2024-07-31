package api

import (
	"database/sql"
	"net/http"
	"vercer/service/user"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var (
	router  *mux.Router
	db      *sql.DB
	session *sessions.CookieStore
)

func init() {
	db, _ = sql.Open(
		"postgres",
		"postgres://default:Ut4uNix0wdRk@ep-polished-sea-a1efivnq.ap-southeast-1.aws.neon.tech:5432/verceldb?sslmode=require",
	)

	session = sessions.NewCookieStore([]byte("514701"))

	router = mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	user.RegisterRoutes(v1, session)
	user.RegisterStore(db)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://go-react-api-web.vercel.app"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"Accept",
			"Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler := c.Handler(router)
	http.Handle("/", handler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
