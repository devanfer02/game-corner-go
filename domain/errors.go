package domain

import (
	"errors"
	"net/http"
)

var (
	ErrServerError 	= errors.New("internal server error")
	ErrNotFound		= errors.New("requested item not found")
	ErrConflict		= errors.New("item already exist")
	ErrBadRequest 	= errors.New("given body request is not valid")
	ErrBadParam		= errors.New("given param request is not valid")
)

func GetErrorCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err  {
		case ErrServerError :
			return http.StatusInternalServerError
		case ErrNotFound : 
			return http.StatusNotFound 
		case ErrConflict :
			return http.StatusConflict
		case ErrBadParam :
			return http.StatusBadRequest
		case ErrBadRequest : 
			return http.StatusBadRequest
		default : 
			return http.StatusInternalServerError
	}
}