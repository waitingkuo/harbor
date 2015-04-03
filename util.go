package main

import (
	"io"
	"net/http"
	"os"
)

func Download(link string, dstPath string) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	f, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)

	return err
}
