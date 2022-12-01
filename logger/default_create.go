package logger

import (
	"hxsg/config"
	"log"
	"os"
	"strings"
)

func NewDefaultLogger(cfg config.Configuration) {
	// 使用 Configuration
	var level LogLevel = LevelDebug
	if configLevelString, found := cfg.GetString("logging:level"); found {
		level = LogLevelFromString(configLevelString)
	}

	flags := log.Lmsgprefix | log.Ltime
	DefaultLog = &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			LevelTrace:       log.New(os.Stdout, "TRACE ", flags),
			LevelDebug:       log.New(os.Stdout, "DEBUG ", flags),
			LevelInformation: log.New(os.Stdout, "INFO ", flags),
			LevelWarning:     log.New(os.Stdout, "WARNING ", flags),
			LevelFatal:       log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}

var DefaultLog *DefaultLogger

func Init() {
	NewDefaultLogger(config.Cfg)
}

func Close() {
	//关闭所有打开的Log文件
}

func LogLevelFromString(val string) (level LogLevel) {
	switch strings.ToLower(val) {
	case "debug":
		level = LevelDebug
	case "information":
		level = LevelInformation
	case "warning":
		level = LevelWarning
	case "fatal":
		level = LevelFatal
	case "none":
		level = LevelNone
	}
	return
}
