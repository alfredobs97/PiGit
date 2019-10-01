package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createRepo/{dir}", createRepo)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
