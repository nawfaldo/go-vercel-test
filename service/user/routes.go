package user

import (
	"database/sql"
	"net/http"
	"vercer/utils"

	"github.com/gorilla/mux"
)

var db *sql.DB

func RegisterRoutes(router *mux.Router, database *sql.DB) {
	db = database

	router.HandleFunc("/users", handleGetUsers).Methods("GET")
	router.HandleFunc("/user", handleCreateUsers).Methods("POST")
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := GetUsers(db)

	utils.WriteJSON(w, http.StatusAccepted, users)
}

func handleCreateUsers(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusAccepted, "hello")
}
