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

func ErrAuthResendCodeLater(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100007,
		Message:   "Resend later",
	}
}

func ErrAuthResendCode(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 100008,
		Message:   "Resend code fail",
	}
}

func ErrAuthUnauthorized(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: 100008,
		Message:   "Unauthorized",
	}
}
