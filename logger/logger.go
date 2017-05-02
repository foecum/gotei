package logger

import "github.com/fatih/color"

// Logger ...
type Logger interface {
	Error(msg string)
	Warning(msg string)
	Notice(msg string)
	Success(msg string)
}

// New ...
func New() Logger {
	return &logger{}
}

type logger struct {
	colorMe color.Color
}

var logPrefix = "[Gotei]: "

func (l *logger) Error(msg string) {
	if len(msg) > 0 {
		c := color.New(color.FgRed)
		c.Println(logPrefix + msg)
	}
}

func (l *logger) Warning(msg string) {
	if len(msg) > 0 {
		c := color.New(color.FgYellow)
		c.Println(logPrefix + msg)
	}
}

func (l *logger) Notice(msg string) {
	if len(msg) > 0 {
		c := color.New(color.FgBlue)
		c.Println(logPrefix + msg)
	}
}

func (l *logger) Success(msg string) {
	if len(msg) > 0 {
		c := color.New(color.FgGreen)
		c.Println(logPrefix + msg)
	}
}
