package myerror

import "net/http"

func ErrFileCreateDest(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 108001,
		Message:   "Fail to create destination file",
	}
}

func ErrFileCopy(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 108002,
		Message:   "Fail to copy file",
	}
}

func ErrFileUploadS3(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 108003,
		Message:   "Fail to upload file to S3",
	}
}

func ErrFileOpenFile(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 108004,
		Message:   "Fail to open file",
	}
}

func ErrFileMissingFile(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 108005,
		Message:   "Missing file",
	}
}

func ErrFileMissingMultipartForm(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 108006,
		Message:   "Missing MultipartForm",
	}
}

func ErrFileDetectMimeFile(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 108007,
		Message:   "Fail to detect mime file",
	}
}

func ErrFileSeek(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 108008,
		Message:   "Fail to seek file",
	}
}

func ErrFilePresignS3(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 108009,
		Message:   "Fail to presign",
	}
}
