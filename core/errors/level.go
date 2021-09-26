package errors

import (
	errbase "errors"
)

type level string

const (
	infoLevel  level = "info"
	warnLevel  level = "warn"
	errorLevel level = "error"
	fatalLevel level = "fatal"
)

func (e *AppError) SetInfo() *AppError {
	e.level = infoLevel
	return e
}

func (e *AppError) SetWarn() *AppError {
	e.level = warnLevel
	return e
}

func (e *AppError) SetError() *AppError {
	e.level = errorLevel
	return e
}

func (e *AppError) SetFatal() *AppError {
	e.level = fatalLevel
	return e
}

func (e *AppError) IsInfo() bool  { return e.checkLevel(infoLevel) }
func (e *AppError) IsWarn() bool  { return e.checkLevel(warnLevel) }
func (e *AppError) IsError() bool { return e.checkLevel(errorLevel) }
func (e *AppError) IsFatal() bool { return e.checkLevel(fatalLevel) }

func (e *AppError) checkLevel(lv level) bool {
	if e.level != `` {
		return e.level == lv
	}
	var next *AppError
	if errbase.As(e.err, &next) {
		return next.checkLevel(lv)
	}
	return false
}
