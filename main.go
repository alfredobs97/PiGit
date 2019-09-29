package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createRepo/{dir}", createRepo)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

}

func checkDiskSpace() bool {
	var stat syscall.Statfs_t
	var isSpace = false
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	syscall.Statfs(wd, &stat)

	var freePercentDisk = (float32(stat.Bavail) / float32(stat.Blocks)) * 100

	if freePercentDisk > 10 {
		isSpace = true
	}

	return isSpace
}

func initGit(path string) error {
	cmd := exec.Command("git", "init", "--bare")
	cmd.Dir = path + ".git"

	_, err := cmd.Output()

	return err
}
func createDirectory(path string) error {
	var err error

	if checkDiskSpace() {
		path = "/repos/" + path

		err = os.MkdirAll(path+".git", os.ModePerm)

		if err == nil {
			err = initGit(path)
		}
	} else {
		err = errors.New("No space disk available")
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
