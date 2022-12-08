package errors

import (
	"errors"
	"net/http"

	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/dto"
)

var (
	ErrUnknown                   = errors.New("internal server error")
	ErrInvalidRequestBody        = errors.New("invalid request body")
	ErrNotFound                  = errors.New("data not found")
	ErrUserExists                = errors.New("email is already taken")
	ErrInvalidResources          = errors.New("resources is empty")
	ErrInvalidCred               = errors.New("password is invalid")
	ErrUnauthorized              = errors.New("user is unauthorized")
	ErrDataNotFound              = errors.New("data not found")
	ErrDuplicateEntry            = errors.New("duplicate field entry")
	ErrEmailFormat               = errors.New("wrong email format")
	ErrMismatchedHashAndPassword = errors.New("wrong password")
	ErrOnlyAdmin                 = errors.New("only admin can access")
	ErrTopupBalance              = errors.New("failed to topup balance")
)

func NewErrorData(code int, message string) dto.ErrorData {
	return dto.ErrorData{
		Code:    code,
		Message: message,
	}
}

func GetErrorResponseMetaData(err error) (er dto.ErrorData) {
	er, ok := errorMap[err]
	if !ok {
		er = errorMap[ErrUnknown]
	}
	return
}

var errorMap = map[error]dto.ErrorData{
	ErrUnknown:                   NewErrorData(http.StatusInternalServerError, ErrUnknown.Error()),
	ErrInvalidRequestBody:        NewErrorData(http.StatusBadRequest, ErrInvalidRequestBody.Error()),
	ErrNotFound:                  NewErrorData(http.StatusNotFound, ErrNotFound.Error()),
	ErrUserExists:                NewErrorData(http.StatusBadRequest, ErrUserExists.Error()),
	ErrInvalidResources:          NewErrorData(http.StatusNotFound, ErrInvalidResources.Error()),
	ErrInvalidCred:               NewErrorData(http.StatusBadRequest, ErrInvalidCred.Error()),
	ErrUnauthorized:              NewErrorData(http.StatusUnauthorized, ErrUnauthorized.Error()),
	ErrDataNotFound:              NewErrorData(http.StatusNotFound, ErrDataNotFound.Error()),
	ErrDuplicateEntry:            NewErrorData(http.StatusBadRequest, ErrDuplicateEntry.Error()),
	ErrEmailFormat:               NewErrorData(http.StatusBadRequest, ErrEmailFormat.Error()),
	ErrMismatchedHashAndPassword: NewErrorData(http.StatusBadRequest, ErrMismatchedHashAndPassword.Error()),
	ErrOnlyAdmin:                 NewErrorData(http.StatusForbidden, ErrOnlyAdmin.Error()),
	ErrTopupBalance:              NewErrorData(http.StatusBadRequest, ErrTopupBalance.Error()),
}
