package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createRepo/{dir}", createRepo)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
