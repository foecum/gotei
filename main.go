package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/foecum/gotei2.0/logger"
	"github.com/foecum/gotei2.0/templates"
	"github.com/urfave/cli"
)

type appContent struct {
	AppName string
	Port    int
}

var log = logger.New()

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"-n"},
			Usage:   "A live reload tool for development web apps written in go. Create a new web app project",
			Action: func(c *cli.Context) error {
				projectName := c.Args().First()
				if len(projectName) < 1 {
					log.Error("argument [appname|controller|model] is missing")
					return nil
				}
				projectPath := os.Getenv("GOPATH") + "/src/" + projectName
				_, err := os.Stat(projectPath)
				if err != nil {
					if os.IsNotExist(err) {
						fmt.Printf("Creating %s\n", projectPath)
						os.MkdirAll(projectPath, os.ModePerm)

						if err = createProjectStructure(projectPath, appContent{AppName: projectName, Port: 8080}); err != nil {
							os.RemoveAll(projectPath)
							log.Error("An error occurred while creating your project")
							log.Error(err.Error())
							return nil
						}
						log.Success("Your new Application was created")
						return nil
					}
				}

				log.Warning(fmt.Sprintf("Directory %s already exists", projectPath))
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:  "appname",
					Usage: "add a new project",
					Action: func(c *cli.Context) error {
						fmt.Println(os.Getenv("GOPATH"))
						return nil
					},
				},
				{
					Name:  "controller",
					Usage: "add a new controller to the project",
					Action: func(c *cli.Context) error {
						fmt.Println(os.Getenv("GOPATH"))
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

func createProjectStructure(path string, projectMeta appContent) error {
	folders := []string{"controllers", "models", "routers", "public", "views", "tests"}
	templatesContent := templates.GetTemplateContent()

	for i := range folders {
		folderPath := path + string(os.PathSeparator) + folders[i]
		_, err := os.Stat(folderPath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Success(fmt.Sprintf("Create %s", folderPath))
				os.MkdirAll(folderPath, os.ModePerm)
				if err = createTemplateFiles(folderPath+string(os.PathSeparator)+folders[i]+".go", templatesContent[folders[i]], projectMeta); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	if err := createTemplateFiles(path+string(os.PathSeparator)+"main.go", templatesContent["main"], projectMeta); err != nil {
		return err
	}
	return nil
}

func createTemplateFiles(fileName, templateContent string, appName appContent) error {

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	var parsedContent bytes.Buffer
	if appName.AppName != "" {
		t := template.New("temp")

		t, err = t.Parse(templateContent)
		if err != nil {
			return err
		}

		t.Execute(&parsedContent, appName)
		templateContent = parsedContent.String()
	}

	_, err = f.Write([]byte(templateContent))
	if err != nil {
		return err
	}
	f.Sync()

	return nil
}
