package builder

import (
	"errors"
	"os/exec"
	"runtime"
	"strings"

	"github.com/foecum/gotei/logger"
)

//Builder ...
type Builder interface {
	Build() error
}

var log = logger.New()

//c.GlobalString("path"), c.GlobalString("bin"), c.GlobalBool("godep")
type builder struct {
	path         string
	name         string
	dependencies bool
}

//New ... creates a new builder
func New(path string, name string, dependencies bool) Builder {
	// does not work on Windows without the ".exe" extension
	if runtime.GOOS == "windows" {
		if !strings.HasSuffix(name, ".exe") { // check if it already has the .exe extension
			name += ".exe"
		}
	}

	return &builder{path: path, name: name, dependencies: dependencies}
}

func (b *builder) Build() error {
	var cmd *exec.Cmd
	if b.dependencies {
		cmd = exec.Command("godep", "go", "build", "-o", b.name)
	} else {
		cmd = exec.Command("go", "build", "-o", b.name)
	}
	cmd.Dir = b.path

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	if !cmd.ProcessState.Success() {
		return errors.New(string(output))
	}
	log.Notice("Build was successful.")

	return nil
}
