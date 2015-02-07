package logging

import (
	"log"
)

type Logger struct {
	ApplicationName string `json:"applicationName"`
}

func NewLogger(applicationName string) *Logger {
	return &Logger{
		ApplicationName: applicationName,
	}
}

func (this *Logger) writeLog(logType LogType, message ...interface{}) {
	log.SetPrefix(this.ApplicationName + ": " + logType.String() + " - ")

	for _, item := range message {
		log.Print(item)
	}
}

func (this *Logger) Info(message ...interface{}) {
	this.writeLog(INFO, message)
}

func (this *Logger) Warning(message ...interface{}) {
	this.writeLog(WARN, message)
}

func (this *Logger) Error(message ...interface{}) {
	this.writeLog(ERROR, message)
}

func (this *Logger) Fatal(message ...interface{}) {
	this.writeLog(FATAL, message)
}
