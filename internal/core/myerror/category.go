package myerror

import "net/http"

func ErrCategoryGet(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 103001,
		Message:   "Get category fail",
	}
}

func ErrCategoryExisted(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 103002,
		Message:   "Category existed",
	}
}

func ErrCategoryCreate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 103003,
		Message:   "Create category fail",
	}
}

func ErrCategoryUpdate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 103004,
		Message:   "Update category fail",
	}
}

func ErrCategoryNotFound(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 103005,
		Message:   "Category not found",
	}
}

func ErrCategoryDataInvalid(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 103006,
		Message:   "Data invalid",
	}
}

func ErrCategoryInvalidParamID(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 103007,
		Message:   "Invalid param id",
	}
}

func ErrCategoryDelete(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 103008,
		Message:   "Delete category",
	}
}
