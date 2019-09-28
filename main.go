package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createRepo/{dir}", createRepo)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

}
func initGit(path string) error {
	cmd := exec.Command("git", "init", "--bare")
	cmd.Dir = path + ".git"

	_, err := cmd.Output()

	return err
}
func createDirectory(path string) error {
	path = "/repos/" + path

	err := os.MkdirAll(path+".git", os.ModePerm)

	if err == nil {
		err = initGit(path)
	}

	return err
}

func createRepo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dir := params["dir"]

	err := createDirectory(dir)

	if err == nil {
		w.WriteHeader(200)
		fmt.Fprintf(w, "<h1> OK </h1>")
	} else {
		w.WriteHeader(503)
	}
}
