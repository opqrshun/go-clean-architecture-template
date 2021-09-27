package controller

type Context interface {
	Param(string) string
	MustGet(string) interface{}
	Bind(interface{}) error
	ShouldBindHeader(interface{}) error
	ShouldBindQuery(interface{}) error
	Status(int)
	JSON(int, interface{})
}

type Logger interface {
	Errorf(format string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}
