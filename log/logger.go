package log

var defaultLogger Logger

type Logger interface {
	// Debug 打印调试日志
	Debug(a ...interface{})
	// Debugf 打印调试模板日志
	Debugf(format string, a ...interface{})
	// Info 打印信息日志
	Info(a ...interface{})
	// Infof 打印信息模板日志
	Infof(format string, a ...interface{})
	// Warn 打印警告日志
	Warn(a ...interface{})
	// Warnf 打印警告模板日志
	Warnf(format string, a ...interface{})
	// Error 打印错误日志
	Error(a ...interface{})
	// Errorf 打印错误模板日志
	Errorf(format string, a ...interface{})
	// Fatal 打印致命错误日志
	Fatal(a ...interface{})
	// Fatalf 打印致命错误模板日志
	Fatalf(format string, a ...interface{})
	// Panic 打印Panic日志
	Panic(a ...interface{})
	// Panicf 打印Panic模板日志
	Panicf(format string, a ...interface{})
}

func init() {
	SetLogger(NewLogger(
	//WithWriter(log.Writer()),
	))
}

// SetLogger 设置日志记录器
func SetLogger(logger Logger) {
	defaultLogger = logger
}

// GetLogger 获取日志记录器
func GetLogger() Logger {
	return defaultLogger
}

// Debug 打印调试日志
func Debug(a ...interface{}) {
	defaultLogger.Debug(a...)
}

// Debugf 打印调试模板日志
func Debugf(format string, a ...interface{}) {
	defaultLogger.Debugf(format, a...)
}

// Info 打印信息日志
func Info(a ...interface{}) {
	defaultLogger.Info(a...)
}

// Infof 打印信息模板日志
func Infof(format string, a ...interface{}) {
	defaultLogger.Infof(format, a...)
}

// Warn 打印警告日志
func Warn(a ...interface{}) {
	defaultLogger.Warn(a...)
}

// Warnf 打印警告模板日志
func Warnf(format string, a ...interface{}) {
	defaultLogger.Warnf(format, a...)
}

// Error 打印错误日志
func Error(a ...interface{}) {
	defaultLogger.Error(a...)
}

// Errorf 打印错误模板日志
func Errorf(format string, a ...interface{}) {
	defaultLogger.Errorf(format, a...)
}

// Fatal 打印致命错误日志
func Fatal(a ...interface{}) {
	defaultLogger.Fatal(a...)
}

// Fatalf 打印致命错误模板日志
func Fatalf(format string, a ...interface{}) {
	defaultLogger.Fatalf(format, a...)
}

// Panic 打印Panic日志
func Panic(a ...interface{}) {
	defaultLogger.Panic(a...)
}

// Fatalf 打印Panic模板日志
func Panicf(format string, a ...interface{}) {
	defaultLogger.Panicf(format, a...)
}
