package user

import (
	"net/http"
	"vercer/utils"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", handleGetUsers).Methods("GET")
	router.HandleFunc("/user", handleCreateUsers).Methods("POST")
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := GetUsers()

	utils.WriteJSON(w, http.StatusAccepted, users)
}

func handleCreateUsers(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusAccepted, "hello")
}
