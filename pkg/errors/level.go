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

func (e *AppError) IsInfo() bool  {
	return e.isLevel(infoLevel)
}
func (e *AppError) IsWarn() bool  {
	 return e.isLevel(warnLevel)
	}
func (e *AppError) IsError() bool { 
	return e.isLevel(errorLevel)
}
func (e *AppError) IsFatal() bool {
	return e.isLevel(fatalLevel)
}

func (e *AppError) isLevel(lv level) bool {
	return e.level == lv
}
