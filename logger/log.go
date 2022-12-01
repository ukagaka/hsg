package logger

func Trace(args ...interface{}) {
	DefaultLog.Trace(args)
}

func Tracef(template string, vals ...interface{}) {
	DefaultLog.Tracef(template, vals...)
}

func Debug(args ...interface{}) {
	DefaultLog.Debug(args)
}

func Debugf(template string, vals ...interface{}) {
	DefaultLog.Debugf(template, vals...)
}

func Info(args ...interface{}) {
	DefaultLog.Info(args)
}

func Infof(template string, vals ...interface{}) {
	DefaultLog.Infof(template, vals...)
}

func Warn(args ...interface{}) {
	DefaultLog.Warn(args)
}

func Warnf(template string, vals ...interface{}) {
	DefaultLog.Warnf(template, vals...)
}

func Error(args ...interface{}) {
	DefaultLog.Error(args)
}

func Errorf(template string, vals ...interface{}) {
	DefaultLog.Errorf(template, vals...)
}

func Panic(args ...interface{}) {
	DefaultLog.Panic(args)
}

func Panicf(template string, vals ...interface{}) {
	DefaultLog.Panicf(template, vals...)
}
