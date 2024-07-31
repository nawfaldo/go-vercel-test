package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello</h1>")
}

func Main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Handler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
