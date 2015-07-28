package logging

import (
	"log"
)

/*
Logger represents an instance of a log structure
*/
type Logger struct {
	ApplicationName string
	LogLevel        LogType

	logLevelInt int
}

/*
Debug writes a debug entry to the log
*/
func (logger *Logger) Debug(message ...interface{}) {
	logger.writeLog(DEBUG, message)
}

/*
Error writes an error message to the log
*/
func (logger *Logger) Error(message ...interface{}) {
	logger.writeLog(ERROR, message)
}

/*
Fatal writes a fatal log messge
*/
func (logger *Logger) Fatal(message ...interface{}) {
	logger.writeLog(FATAL, message)
}

/*
Info writes an information message to the log
*/
func (logger *Logger) Info(message ...interface{}) {
	logger.writeLog(INFO, message)
}

/*
NewLogger creates a new Logger instance with a specific application name.
The application name is a prefix used in all log messges.
*/
func NewLogger(applicationName string) *Logger {
	return &Logger{
		ApplicationName: applicationName,
		LogLevel:        DEBUG,

		logLevelInt: int(DEBUG),
	}
}

/*
NewLoggerWithMinimumLevel creates a new Logger instance that only logs
messages with a specified log type level or higher.
*/
func newLoggerWithMinimumLevel(applicationName string, logLevel LogType) *Logger {
	return &Logger{
		ApplicationName: applicationName,
		LogLevel:        logLevel,

		logLevelInt: int(logLevel),
	}
}

/*
Warning writes a warning message to the log
*/
func (logger *Logger) Warning(message ...interface{}) {
	logger.writeLog(WARN, message)
}

func (logger *Logger) writeLog(logType LogType, message ...interface{}) {
	logLevelInt := int(logType)

	if logLevelInt >= logger.logLevelInt {
		log.SetPrefix(logger.ApplicationName + ": " + logType.String() + " - ")

		for _, item := range message {
			log.Print(item)
		}
	}
}
