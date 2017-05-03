package runner

import "testing"

var testEngine Runner

func TestErrorWithFalse(t *testing.T) {
	testEngine = New("path/to/dir", "temp app", []string{"arg1", "arg2", "arg2"})
}
