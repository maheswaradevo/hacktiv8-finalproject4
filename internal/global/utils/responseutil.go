package utils

import (
	"net/http"

	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
)

func NewSuccessResponseWriter(rw http.ResponseWriter, status string, code int, data interface{}) dto.BaseResponse {
	return BaseResponseWriter(rw, code, status, nil, data)
}

func NewErrorResponse(rw http.ResponseWriter, err error) dto.ErrorData {
	errMap := errors.GetErrorResponseMetaData(err)
	return dto.ErrorData{
		Code:    errMap.Code,
		Message: errMap.Message,
	}
}

func BaseResponseWriter(rw http.ResponseWriter, code int, status string, er *dto.ErrorData, data interface{}) dto.BaseResponse {
	res := dto.BaseResponse{
		Status: status,
		Data:   data,
		Error:  er,
	}
	return res
}
