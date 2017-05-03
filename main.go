package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	"github.com/foecum/gotei/builder"
	"github.com/foecum/gotei/runner"
	"github.com/foecum/gotei2.0/logger"
	"github.com/foecum/gotei2.0/templates"
	"github.com/urfave/cli"
)

type appContent struct {
	AppName string
	Port    int
}

var (
	pkgName string
	path    string
)

var log = logger.New()

func initApp() {
	pkgName, path = "", ""
	cmd := exec.Command("go", "list", "./...")
	buf, err := cmd.Output()
	if err != nil {
		log.Error(fmt.Sprintf("%v", err.Error()))
		os.Exit(2)
	}
	list := strings.Split(string(buf), "\n")
	if len(list) > 0 {
		pkgName = list[0]
	}

	path, err = os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		log.Error(err.Error())
		os.Exit(2)
	}
}

func main() {
	initApp()

	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"-n"},
			Usage:   "A live reload tool for development web apps written in go. Create a new web app project",
			Action:  goteiNewAppAction,
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
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "Run the gotei tool in the current working directory",
			Action:    goteiAction,
		},
		{
			Name:      "install",
			ShortName: "i",
			Usage:     "Run the gotei tool in the current working directory",
			Action:    goteiInstallAction,
		},
	}
	app.Run(os.Args)
}

func goteiNewAppAction(c *cli.Context) error {
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
			goteiAction(c)
			return nil
		}
	}

	log.Warning(fmt.Sprintf("Directory %s already exists", projectPath))
	return nil
}

func goteiInstallAction(c *cli.Context) {
	log.Notice("Builing application.")

	build := builder.New(path, pkgName, false)
	app := runner.New(path, pkgName, c.Args())
	app.Monitor(path, build.Build)
}

func goteiAction(c *cli.Context) {
	log.Notice("Building application.")
	build := builder.New(path, pkgName, false)
	err := build.Build()
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Success("Build run successfully")
	app := runner.New(path, pkgName, c.Args())
	app.Run()

	log.Notice("Exiting application.")
}

func createProjectStructure(path string, projectMeta appContent) error {
	folders := []string{"controllers", "models", "routers"}
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
