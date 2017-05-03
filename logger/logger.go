package logger

import "github.com/fatih/color"

// Logger ...
type Logger interface {
	Error(msg string) bool
	Warning(msg string) bool
	Notice(msg string) bool
	Success(msg string) bool
}

// New ...
func New() Logger {
	return &logger{}
}

type logger struct {
	colorMe color.Color
}

var logPrefix = "[Gotei]▶ "

func (l *logger) Error(msg string) bool {
	return message(msg, color.FgRed)
}

func (l *logger) Warning(msg string) bool {
	return message(msg, color.FgYellow)
}

func (l *logger) Notice(msg string) bool {
	return message(msg, color.FgBlue)
}

func (l *logger) Success(msg string) bool {
	return message(msg, color.FgGreen)
}

func message(msg string, clr color.Attribute) bool {
	if len(msg) > 0 {
		c := color.New(clr)
		c.Println(logPrefix + msg)
		return true
	}
	return false
}
