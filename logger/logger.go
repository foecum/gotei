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

var logPrefix = "[Gotei]▶ "

func (l *logger) Error(msg string) {
	message(msg, color.FgRed)
}

func (l *logger) Warning(msg string) {
	message(msg, color.FgYellow)
}

func (l *logger) Notice(msg string) {
	message(msg, color.FgBlue)
}

func (l *logger) Success(msg string) {
	message(msg, color.FgGreen)
}

func message(msg string, clr color.Attribute) {
	if len(msg) > 0 {
		c := color.New(clr)
		c.Println(logPrefix + msg)
	}
}
