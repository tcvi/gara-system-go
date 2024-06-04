package myerror

import "net/http"

func ErrUserNotFound(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: 103001,
		Message:   "User not found",
	}
}

func ErrCreateUser(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 103002,
		Message:   "Create user fail",
	}
}

func ErrInvalidRegister(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 103003,
		Message:   "Invalid data",
	}
}

func ErrUserExisted(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 103004,
		Message:   err.Error(),
	}
}

func ErrCreateUserHashPassword(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 103002,
		Message:   "Hash password fail",
	}
}
