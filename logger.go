package sqsworker

import (
	"fmt"
	"log"
)

// Logger interface
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
}

type logger struct {
}

func (l *logger) Debug(i ...interface{}) {
	log.Printf("[DEBUG] %s", fmt.Sprintln(i...))
}

func (l *logger) Info(i ...interface{}) {
	log.Printf("[INFO] %s", fmt.Sprintln(i...))
}

func (l *logger) Error(i ...interface{}) {
	log.Printf("[ERROR] %s", fmt.Sprintln(i...))
}
