package logging

type LogType int

const (
	NONE LogType = iota
	INFO
	WARN
	ERROR
	FATAL
)

var logTypeNames = map[LogType]string{
	NONE:  "None",
	INFO:  "INFO",
	WARN:  "WARNING",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

func (this LogType) String() string {
	return logTypeNames[this]
}
