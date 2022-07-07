package customErrors

import (
	"errors"
	"net/http"
)

var (
	ErrBadInputData = errors.New("bad input data")

	ErrLinkNotFound     = errors.New("link doensn`t exists")
	ErrLinkAlreadyExist = errors.New("link already exist")

	ErrLinkCouldNotBeCreated = errors.New("link could not be created, try again later")
)

var errorToCode = map[error]int{
	ErrBadInputData: http.StatusBadRequest,

	ErrLinkNotFound:          http.StatusNotFound,
	ErrLinkAlreadyExist:      http.StatusConflict,
	ErrLinkCouldNotBeCreated: http.StatusBadRequest,
}

func ConvertErrorToCode(err error) (code int) {
	code, isErrorExist := errorToCode[err]
	if !isErrorExist {
		code = http.StatusInternalServerError
	}
	return
}
