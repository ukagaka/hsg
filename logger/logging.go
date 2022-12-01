package logger

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInformation
	LevelWarning
	LevelFatal
	LevelErrot
	LevelNone
)

type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})

	Debug(string)
	Debugf(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Warn(string)
	Warnf(string, ...interface{})

	Error(string)
	Errorf(string, ...interface{})

	Panic(string)
	Panicf(string, ...interface{})
}
