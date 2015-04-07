package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type HarborStore struct {
	RootPath string
	BinPath  string
}

func NewHarborStore() *HarborStore {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// default path
	rootPath := filepath.Join(usr.HomeDir, ".harbor")
	binPath := filepath.Join(rootPath, "bin")

	err = os.MkdirAll(binPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return &HarborStore{
		RootPath: rootPath,
		BinPath:  binPath,
	}

}
