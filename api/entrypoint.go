package api

import (
	"net/http"
	"vercer/utils"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", hello).Methods("GET")
}

func hello(w http.ResponseWriter, r *http.Request) {
	hello := map[string]string{
		"msg": "hello",
	}

	utils.WriteJSON(w, http.StatusAccepted, hello)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
