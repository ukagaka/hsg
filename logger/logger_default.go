package logger

import (
	"fmt"
	"log"
)

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {

		l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLogger) Trace(args ...interface{}) {
	l.write(LevelTrace, fmt.Sprint(args...))
}

func (l *DefaultLogger) Tracef(template string, args ...interface{}) {
	l.write(LevelTrace, fmt.Sprintf(template, args...))
}

func (l *DefaultLogger) Debug(args ...interface{}) {
	l.write(LevelDebug, fmt.Sprint(args...))
}

func (l *DefaultLogger) Debugf(template string, vals ...interface{}) {
	l.write(LevelDebug, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Info(args ...interface{}) {
	l.write(LevelInformation, fmt.Sprint(args...))
}

func (l *DefaultLogger) Infof(template string, vals ...interface{}) {
	l.write(LevelInformation, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Warn(args ...interface{}) {
	l.write(LevelWarning, fmt.Sprint(args...))
}

func (l *DefaultLogger) Warnf(template string, vals ...interface{}) {
	l.write(LevelWarning, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Error(args ...interface{}) {
	l.write(LevelErrot, fmt.Sprint(args...))
}

func (l *DefaultLogger) Errorf(template string, vals ...interface{}) {
	l.write(LevelErrot, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Panic(args ...interface{}) {
	l.write(LevelFatal, fmt.Sprint(args...))
	if l.triggerPanic {
		panic(args)
	}
}

func (l *DefaultLogger) Panicf(template string, vals ...interface{}) {
	formattedMsg := fmt.Sprintf(template, vals...)
	l.write(LevelFatal, formattedMsg)
	if l.triggerPanic {
		panic(formattedMsg)
	}
}
