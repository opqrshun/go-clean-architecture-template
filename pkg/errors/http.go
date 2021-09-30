package errors

import (
	// "fmt"
	"net/http"
)

// Event
type Event struct {
	code       string
	level      level
	httpStatus int
	displayMsg string
}

var (
	invalidArgValue = Event{
		"invalid_arg_value",
		errorLevel,
		http.StatusBadRequest,
		"requested parameter value is invalid",
	}

	unauthorized = Event{
		"unauthorized",
		errorLevel,
		http.StatusUnauthorized,
		"token is invalid",
	}

	databaseError = Event{
		"database_error",
		errorLevel,
		http.StatusInternalServerError,
		"Something went wrong. please try again",
	}

	//バリデート失敗時、
	// リクエストパラメータが不正のとき。
	invalidRequest = Event{
		"invalid_request",
		errorLevel,
		http.StatusBadRequest,
		"invalid request",
	}

	internal = Event{
		"internal",
		errorLevel,
		http.StatusInternalServerError,
		"internal",
	}

	unknown = Event{
		"unknown",
		errorLevel,
		http.StatusInternalServerError,
		"unknown",
	}

	//usernameConflict info
	usernameConflict = Event{
		"username_conflict",
		infoLevel,
		http.StatusConflict,
		"username is already used, please use another user name",
	}

	//notFound info
	notFound = Event{
		"not_found",
		infoLevel,
		http.StatusNotFound,
		"not found",
	}
)

func (e *AppError) InvalidArgValue() *AppError {
	return e.setEvent(invalidArgValue)
}

func (e *AppError) InvalidRequest() *AppError {
	return e.setEvent(invalidRequest)
}

func (e *AppError) Unauthorized() *AppError {
	return e.setEvent(unauthorized)
}

func (e *AppError) DatabaseError() *AppError {
	return e.setEvent(databaseError)
}

func (e *AppError) UsernameConflict() *AppError {
	return e.setEvent(usernameConflict)
}

func (e *AppError) NotFound() *AppError {
	return e.setEvent(notFound)
}

func (e *AppError) Internal() *AppError {
	return e.setEvent(internal)
}

func (e *AppError) Unknown() *AppError {
	return e.setEvent(unknown)
}

func (e *AppError) setEvent(event Event) *AppError {
	e.code = event.code
	e.level = event.level
	e.httpStatus = event.httpStatus
	e.displayMsg = event.displayMsg
	return e
}

//IsNotFound
//check err, example, an error wrapped with a not found record error
func IsNotFound(err error) bool {
	return checkCode(err, NotFound)
}

func checkCode(err error, code string) bool {
	aerr := AsAppError(err)
	if aerr == nil {
		return false
	}
	return aerr.code == code
}

const (
	InvalidRequest = "invalid_request"
	Unauthorized   = "unauthorized"
	NotFound       = "not_found"
	Internal       = "internal_error"
	Forbidden      = "forbidden"
	Unknown        = "unknown"
)
