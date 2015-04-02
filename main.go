package main

import (
	"fmt"
	"github.com/codegangsta/cli"
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

func main() {
	store := NewHarborStore()

	app := cli.NewApp()
	app.Name = "harbor"
	app.Commands = []cli.Command{

		{
			Name:  "env",
			Usage: "Display command for ...",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "install",
			Usage: "install package",
			Action: func(c *cli.Context) {
				packageName := c.Args().First()
				if packageName == "compose" || packageName == "docker-compose" {

					fmt.Println("Installing docker-compose ...")

					version := "1.2.0rc3"

					sysName, err := GetSysName()
					if err != nil {
						log.Fatal(err)
					}
					if sysName != "Linux" && sysName != "Darwin" {
						fmt.Printf("System %s is not supported\n", sysName)
						return
					}

					machine, err := GetMachine()
					if err != nil {
						log.Fatal(err)
					}
					if machine != "x86_64" {
						fmt.Printf("Machine %s is not supported\n", machine)
						return
					}

					link := fmt.Sprintf("https://github.com/docker/compose/releases/download/%s/docker-compose-%s-%s", version, sysName, machine)

					dstPath := filepath.Join(store.BinPath, "docker-compose")
					Download(link, dstPath)

				} else if packageName == "machine" || packageName == "docker-machine" {
					fmt.Println("Installing docker-machine ...")

					version := "v0.2.0-rc3"

					sysName, err := GetSysName()
					if err != nil {
						log.Fatal(err)
					}
					if sysName == "Darwin" {
						sysName = "darwin"
					} else if sysName == "Linux" {
						sysName = "linux"
					} else if sysName == "Windows" {
						sysName = "windows"
					} else {
						fmt.Printf("System %s is not supported\n", sysName)
						return
					}

					machine, err := GetMachine()
					if err != nil {
						log.Fatal(err)
					}
					if machine == "x86_64" {
						machine = "amd64"
					} else if machine == "i386" || machine == "i686" {
						machine = "386"
					} else {
						fmt.Printf("Machine %s is not supported\n", machine)
						return
					}

					var extenion string
					if machine == "window" {
						extenion = ".exe"
					}

					link := fmt.Sprintf("https://github.com/docker/machine/releases/download/%s/docker-machine_%s-%s%s", version, sysName, machine, extenion)
					dstPath := filepath.Join(store.BinPath, "docker-machine")
					Download(link, dstPath)

				}
			},
		},
	}
	app.Run(os.Args)
}
