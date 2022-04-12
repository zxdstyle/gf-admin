package middleware

import (
	"gf-admin/pkg/responses"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
	"net/http"
)

func Response(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		ctx = r.Context()
		err = r.GetError()
		res = r.GetHandlerResponse()
	)

	if err == nil && r.Response.Status == http.StatusOK {
		r.Response.ClearBuffer()
		r.Response.WriteHeader(http.StatusOK)
		internalErr := r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    res,
		})
		if internalErr != nil {
			g.Log().Errorf(ctx, `%+v`, internalErr)
		}
		return
	}

	var code gcode.Code
	switch r.Response.Status {
	case http.StatusOK:
		if val, ok := err.(*gerror.Error); ok && val.Code().Code() > http.StatusOK {
			code = val.Code()
		} else {
			code = gcode.New(http.StatusInternalServerError, err.Error(), nil)
		}
	case http.StatusNotFound:
		code = responses.NotFound
	case http.StatusForbidden:
		code = responses.Forbidden
	case http.StatusBadRequest:
		code = responses.BadRequest
	case http.StatusUnprocessableEntity:
		code = responses.UnprocessableEntity
	default:
		code = gcode.New(http.StatusInternalServerError, err.Error(), nil)
	}

	if val, ok := err.(gvalid.Error); ok {
		code = gcode.New(http.StatusUnprocessableEntity, val.Error(), nil)
	}

	r.Response.ClearBuffer()
	r.Response.WriteHeader(code.Code())
	internalErr := r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data:    nil,
	})
	if internalErr != nil {
		g.Log().Errorf(ctx, `%+v`, internalErr)
	}
}
