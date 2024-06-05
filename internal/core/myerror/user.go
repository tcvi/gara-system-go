package myerror

import "net/http"

func ErrUserNotFound(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: 101001,
		Message:   "User not found",
	}
}

func ErrCreateUser(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 101002,
		Message:   "Create user fail",
	}
}

func ErrInvalidRegister(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 101003,
		Message:   "Invalid data",
	}
}

func ErrUserExisted(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 101004,
		Message:   err.Error(),
	}
}

func ErrCreateUserHashPassword(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 101005,
		Message:   "Hash password fail",
	}
}

func ErrInvalidLogin(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 101006,
		Message:   "Invalid data",
	}
}
