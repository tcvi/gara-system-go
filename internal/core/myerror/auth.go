package myerror

import "net/http"

func ErrInvalidVerify(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100001,
		Message:   "Invalid data",
	}
}

func ErrAuthNotFoundAccount(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100002,
		Message:   "Not found account",
	}
}

func ErrAuthAccountHasVerified(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100003,
		Message:   "Account has been verified",
	}
}

func ErrAuthExpiredCode(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100004,
		Message:   "Expired code",
	}
}

func ErrAuthWrongCode(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100005,
		Message:   "Wrong code",
	}
}

func ErrAuthUpdateAccount(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100005,
		Message:   "Update account fail",
	}
}

func ErrAuthInvalidDataRequest(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100006,
		Message:   "Invalid data",
	}
}
