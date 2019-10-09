package main

import (
	"errors"
	"log"
	"os"
	"syscall"

	"gopkg.in/src-d/go-git.v4"
)

func initGit(path string) error {
	if checkDiskSpace() {
		_, err := git.PlainInit("repos/"+path, true)

		if err != nil {
			return err
		}
	} else {
		return errors.New("No space disk available")
	}

	return nil
}

func checkDiskSpace() bool {
	var stat syscall.Statfs_t

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	syscall.Statfs(wd, &stat)

	freePercentDisk := (float32(stat.Bavail) / float32(stat.Blocks)) * 100

	if freePercentDisk > 10 {
		return true
	}

	return false
}
