package api

import (
	"net/http"
	"vercer/utils"

	"github.com/gorilla/mux"
)

func API(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods("GET")

	router.ServeHTTP(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {

	hello := map[string]string{
		"msg": "hello",
	}

	utils.WriteJSON(w, http.StatusAccepted, hello)
}
