package log

import (
	"fmt"
	"log"
)

// Logger for logging
type Logger struct{}

// New returns new Log
func New() Logger {
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime)
	return &Logger{}
}

// LogInfo logs with [INFO] prefix
func (l *Logger) LogInfo(v ...interface{}) {
	log.Println("[INFO]", fmt.Sprintln(v...))
}

// LogError logs with [ERROR] prefix
func (l *Logger) LogError(v ...interface{}) {
	log.Println("[ERROR]", fmt.Sprintln(v...))
}
