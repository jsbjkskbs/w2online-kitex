package utils

import (
	"errors"
	"work/pkg/errno"
)

type BaseHttpResponse struct {
	StatusCode int64
	StatusMsg  string
}

func baseHttpResponse(err errno.Errno) *BaseHttpResponse {
	return &BaseHttpResponse{
		StatusCode: err.Code,
		StatusMsg:  err.Message,
	}
}

func CreateBaseHttpResponse(err error) *BaseHttpResponse {
	if err == nil {
		return baseHttpResponse(errno.NoError)
	}

	e := errno.Errno{}
	if errors.As(err, &e) {
		return baseHttpResponse(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseHttpResponse(s)
}
