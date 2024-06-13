package myerror

import (
	"fmt"
	"net/http"
)

func ErrItemInvalidData(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 104001,
		Message:   "Invalid data",
	}
}

func ErrItemCreate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 104002,
		Message:   "Create item fail",
	}
}

func ErrItemParamIDInvalid(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 104003,
		Message:   "Param id invalid",
	}
}

func ErrItemItemNotFound(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 104004,
		Message:   "Item not found",
	}
}

func ErrItemCategoryNotFound(err error, id int64) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 104005,
		Message:   fmt.Sprintf("Category id = %d not found", id),
	}
}

func ErrItemUpdate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 104006,
		Message:   "Update item fail",
	}
}

func ErrItemGetList(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 104007,
		Message:   "Get list item fail",
	}
}

func ErrItemGetCategory(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 104008,
		Message:   "Get category list fail",
	}
}
