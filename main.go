package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
	"runtime"
)

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

					goos := runtime.GOOS
					if goos == "linux" {
						goos = "Linux"
					} else if goos == "darwin" {
						goos = "Darwin"
					} else {
						fmt.Printf("OS %s is not supported\n", goos)
						return
					}

					goarch := runtime.GOARCH
					if goarch != "x86_64" {
						goarch = "amd64"
					} else {
						fmt.Printf("Machine %s is not supported\n", goarch)
						return
					}

					link := fmt.Sprintf("https://github.com/docker/compose/releases/download/%s/docker-compose-%s-%s", version, goos, goarch)

					dstPath := filepath.Join(store.BinPath, "docker-compose")
					Download(link, dstPath)

				} else if packageName == "machine" || packageName == "docker-machine" {
					fmt.Println("Installing docker-machine ...")

					version := "v0.2.0-rc3"

					goos := runtime.GOOS
					if goos != "linux" && goos != "windows" && goos != "darwin" {
						fmt.Printf("OS %s is not supported\n", goos)
						return
					}

					goarch := runtime.GOARCH
					if goarch != "amd64" && goarch != "386" {
						fmt.Printf("OS %s is not supported\n", goarch)
						return
					}

					extenion := ""
					if goos == "windows" {
						extenion = ".exe"
					}

					link := fmt.Sprintf("https://github.com/docker/machine/releases/download/%s/docker-machine_%s-%s%s", version, goos, goarch, extenion)
					dstPath := filepath.Join(store.BinPath, "docker-machine")
					Download(link, dstPath)

				}
			},
		},
	}
	app.Run(os.Args)
}
