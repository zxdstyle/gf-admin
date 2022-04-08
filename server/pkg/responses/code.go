package responses

import "net/http"

type restfulCode struct {
	code    int
	message string
	data    interface{}
}

var (
	NotFound            = newCode(http.StatusNotFound, nil)
	Forbidden           = newCode(http.StatusForbidden, nil)
	BadRequest          = newCode(http.StatusBadRequest, nil)
	UnprocessableEntity = newCode(http.StatusUnprocessableEntity, nil)
	InternalError       = newCode(http.StatusInternalServerError, nil)
)

func newCode(code int, data interface{}) restfulCode {
	return restfulCode{
		code: code,
		data: data,
	}
}

func (r restfulCode) Code() int {
	return r.code
}

func (r restfulCode) Message() string {
	return http.StatusText(r.code)
}

func (r restfulCode) Detail() interface{} {
	return r.data
}
