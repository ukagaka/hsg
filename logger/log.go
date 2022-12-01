package logger

func Trace(template string, vals ...interface{}) {
	DefaultLog.Tracef(template, vals...)
}

func Debug(template string, vals ...interface{}) {
	DefaultLog.Debugf(template, vals...)
}

func Info(template string, vals ...interface{}) {
	DefaultLog.Infof(template, vals...)
}

func Warn(template string, vals ...interface{}) {
	DefaultLog.Warnf(template, vals...)
}

func Error(template string, vals ...interface{}) {
	DefaultLog.Errorf(template, vals...)
}

func Panic(template string, vals ...interface{}) {
	DefaultLog.Panicf(template, vals...)
}
