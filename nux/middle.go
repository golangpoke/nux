package nux

import (
	"fmt"
	"github.com/golangpoke/nux/nlog"
)

func Logger() Middleware {
	return func(next HandleFunc) HandleFunc {
		return func(req *Request) Response {
			nlog.INFOf(fmt.Sprintf("%s %s", req.Request().Method, req.Request().URL.Path))
			return next(req)
		}
	}
}

func Recovery() Middleware {
	return func(next HandleFunc) HandleFunc {
		return func(req *Request) Response {
			defer nlog.Recovery()
			return next(req)
		}
	}
}