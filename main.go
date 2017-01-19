package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Create a new web app project",
			Action: func(c *cli.Context) error {
				projectName := c.Args().First()
				if len(projectName) < 1 {
					log.Println("argument <appname> is missing")
					return nil
				}
				projectPath := os.Getenv("GOPATH") + "/src/" + projectName
				_, err := os.Stat(projectPath)
				if err != nil {
					if os.IsNotExist(err) {
						fmt.Printf("Create %s\n", projectPath)
						os.MkdirAll(projectPath, os.ModePerm)

						createProjectStructure(projectPath)
						fmt.Printf("Your new Application was created\n")
						return nil
					}
				}

				fmt.Printf("Directory %s already exists\n", projectPath)
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:  "controller",
					Usage: "add a new controller to the project",
					Action: func(c *cli.Context) error {
						fmt.Println(os.Getenv("GOPATH"))
						//fmt.Println("new controller added: ", c.Args().First(), " ", os.Getenv("GOPATH"))
						return nil
					},
				},
				{
					Name:  "model",
					Usage: "add a new model to the project",
					Action: func(c *cli.Context) error {
						fmt.Println("new model added: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func createProjectStructure(path string) {
	folders := []string{"controllers", "models", "routers", "config", "public", "views", "tests"}

	for i := range folders {
		folderPath := path + "/" + folders[i]
		_, err := os.Stat(folderPath)
		if err == nil {
			os.RemoveAll(folderPath)
		}
		if os.IsNotExist(err) {
			fmt.Printf("Create %s\n", folderPath)
			os.MkdirAll(folderPath, os.ModePerm)
		}
	}
}
