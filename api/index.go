package handler

import (
	"fmt"
	"net/http"
	"vercer/utils"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello</h1>")
}

func Main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", hello).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	hello := map[string]string{
		"msg": "hello",
	}

	utils.WriteJSON(w, http.StatusAccepted, hello)
}
