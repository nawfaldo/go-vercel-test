package api

import (
	"database/sql"
	"net/http"
	"vercer/service/user"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
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
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://go-react-api-web.vercel.app")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	corsHandler := cors(router)

	corsHandler.ServeHTTP(w, r)
}
