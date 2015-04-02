package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func GetSysName() (string, error) {

	cmd := exec.Command("uname", "-s")
	sysNameByte, err := cmd.Output()
	if err != nil {
		return "", err
	}

	sysName := strings.Trim(string(sysNameByte), "\n")

	return sysName, err
}
func GetMachine() (string, error) {

	cmd := exec.Command("uname", "-m")
	machineByte, err := cmd.Output()
	if err != nil {
		return "", err
	}

	machine := strings.Trim(string(machineByte), "\n")

	return machine, err
}
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
