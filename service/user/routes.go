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
	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type User struct {
		Name string `json:"name"`
	}

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	utils.WriteJSON(w, http.StatusAccepted, users)
}

func handleCreateUsers(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusAccepted, "hello")
}
