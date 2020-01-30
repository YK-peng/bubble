package xlog

type Xlogger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})

	//Warning(args ...interface{})
	//Warningf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}
