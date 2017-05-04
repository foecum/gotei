package runner

import (
	"fmt"
	"testing"
)

var testEngine Runner

// SystemCallsImplementationTestFail ...
type SystemCallsImplementationTestFail struct {
}

func (SystemCallsImplementationTestFail) Getwd() (string, error) {
	return "", fmt.Errorf("An error occured")
}

func TestErrorWithFalse(t *testing.T) {
	testEngine = New("path/to/dir", "temp app", []string{"arg1", "arg2", "arg2"})

	testFail := SystemCallsImplementationTestFail{}
	if testEngine.Run(testFail) != false {
		t.Error("Expected false but got something else")
	}
}
