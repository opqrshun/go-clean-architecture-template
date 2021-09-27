package errors

import (
	errbase "errors"
	"fmt"
	"net/http"
)

type AppError struct {
	err error
	msg string

	// Error level
	// When loggint
	level level

	//for Respond API
	code       string
	displayMsg string
	httpStatus int
}

// New
func New(msg string) *AppError {
	return create(msg)
}

// Errorf
// not support %w
func Errorf(format string, args ...interface{}) *AppError {
	return create(fmt.Sprintf(format, args...))
}

// Wrapf exec Errorf with wrap
// not support %w
func Wrapf(err error, format string, args ...interface{}) *AppError {
	e := Errorf(format, args...)
	e.err = err
	return e
}

// create  set msg
func create(msg string) *AppError {
	var e AppError
	e.msg = msg
	return &e
}

func Is(err, target error) bool {
	return errbase.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errbase.As(err, target)
}

//Unwrap
func Unwrap(err error) error {
	return errbase.Unwrap(err)
}

func AsAppError(err error) *AppError {
	if err == nil {
		return nil
	}

	var aerr *AppError
	if errbase.As(err, &aerr) {
		return aerr
	}

	return nil
}

//Error  return message
func (e *AppError) Error() string {
	return e.msg
}

//Code return code
func (e *AppError) Code() string {
	if e.code != "" {
		return e.code
	}
	return `not_defined`
}

//HttpStatus return http status
func (e *AppError) HttpStatus() int {
	if e.httpStatus != 0 {
		return e.httpStatus
	}
	return http.StatusInternalServerError
}

//HttpStatus return http status
func (e *AppError) DisplayMessage() string {
	return e.displayMsg
}

//Unwrap
func (e *AppError) Unwrap() error {
	return e.err
}

//Wrap
func (e *AppError) Wrap(err error) *AppError {
	e.err = err
	return e
}
