package nux

import (
	"fmt"
	"github.com/golangpoke/nux/nlog"
)

func Logger() Middleware {
	return func(next HandleFunc) HandleFunc {
		return func(req *Request) Response {
			nlog.INFOf(fmt.Sprintf("%s %s", req.Method(), req.Url()))
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

func CORS() Middleware {
	return func(next HandleFunc) HandleFunc {
		return func(req *Request) Response {
			req.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			req.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			return next(req)
		}
	}
}
