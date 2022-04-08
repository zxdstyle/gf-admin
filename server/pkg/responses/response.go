package responses

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"net/http"
)

type (
	Response struct {
		Code    int
		Message string
		Data    interface{}
	}

	CreateRes interface {
	}

	UpdateRes interface {
	}

	ShowRes interface {
	}

	DestroyRes interface {
	}
)

func Success() *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    nil,
	}
}

func SuccessWithData(data interface{}) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

func Failed(er error) error {
	return gerror.NewCode(BadRequest, er.Error())
}
