package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/foecum/gotei/builder"
	"github.com/foecum/gotei/logger"
	"github.com/foecum/gotei/sockets"
)

var log = logger.New()
var goteiSockets sockets.GoteiSocketActions

// Runner ... Interface with two main functionalities the app implements
type Runner interface {
	Run(sysCall SystemCalls) bool
	Monitor(cwd string, action func() error) error
	Kill() error
}

type engine struct {
	name      string //Name of the app
	cmd       *exec.Cmd
	args      []string
	beginning time.Time
	builder   builder.Builder
	appMeta   map[string]interface{}
}

//New ...created a new instance of engine
func New(path, name string, args []string) Runner {
	return &engine{
		name:      name,
		args:      args,
		beginning: time.Now(),
		builder:   builder.New(path, name, false),
		appMeta:   make(map[string]interface{}),
	}
}

// SystemCalls ...
type SystemCalls interface {
	Getwd() (string, error)
}

// SystemCallsImplementation ...
type SystemCallsImplementation struct {
}

// Getwd ...
func (SystemCallsImplementation) Getwd() (string, error) {
	return os.Getwd()
}

//Run ...runs the app and monitors file modification times
func (e *engine) Run(sysCall SystemCalls) bool {
	cwd, err := sysCall.Getwd()

	if err != nil {
		log.Error(err.Error())
		return false
	}

	err = e.start()
	if err != nil {
		log.Error(err.Error())
		return false
	}

	goteiSockets := new(sockets.GoteiSocketActions)
	goteiSockets.StartReloadServer()

	log.Notice("Listening...")
	e.Monitor(cwd, e.restart)
	return true
}

func (e *engine) Monitor(cwd string, action func() error) error {
	for {
		err := filepath.Walk(cwd, func(path string, f os.FileInfo, err error) error {
			// Ignore the .git folder
			if path == ".git" {
				return filepath.SkipDir
			}
			// ignore hidden files and configs and binaries
			file := string(filepath.Base(path)[0])
			if file == "." && !strings.Contains(file, e.name) {
				return nil
			}
			// check for changes in go files and ignore the tests
			if !strings.Contains(f.Name(), "_test.go") && f.ModTime().After(e.beginning) {
				log.Notice(fmt.Sprintf("Modified: %v", f.Name()))
				err := action()
				if err != nil {
					log.Error(err.Error())
				}
				log.Notice("Watching...")
				e.beginning = time.Now()
			}
			return nil
		})

		if err != nil {
			log.Error(err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

func (e *engine) restart() error {
	if err := e.Kill(); err != nil {
		return fmt.Errorf("%v", err)
	}
	log.Success(fmt.Sprintf("%v stopped successfully.", e.name))

	err := e.builder.Build()
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	if err := e.start(); err != nil {
		return fmt.Errorf("%v", err)
	}

	goteiSockets.SendReload()
	log.Success(fmt.Sprintf("%v restarted successfully.", e.name))
	return nil
}

func (e *engine) start() error {
	log.Notice(fmt.Sprintf("starting %v application", e.name))
	e.cmd = exec.Command("./" + e.name)
	err := e.cmd.Start()
	if err != nil {
		return err
	}
	cmnd := exec.Command("open", "http//localhost:8081")
	cmnd.Start()
	return nil
}

func (e *engine) Kill() error {
	log.Notice(fmt.Sprintf("Stopping %v", e.name))
	if err := e.cmd.Process.Kill(); err != nil {
		return fmt.Errorf("failed to stop %s: %v ", e.name, err)
	}
	return nil
}
