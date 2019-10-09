package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func createRepo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dir := params["dir"]

	err := initGit(dir)

	if err == nil {
		w.WriteHeader(200)
		fmt.Fprintf(w, "<h1> OK </h1>")
	} else {
		w.WriteHeader(503)
		fmt.Println(err)
	}
}
