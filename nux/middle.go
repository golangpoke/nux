package nux

import (
	"fmt"
	"github.com/golangpoke/nux/nlog"
	"net/http"
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
			req.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
			req.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

			if req.r.Method == "OPTIONS" {
				req.Writer.WriteHeader(http.StatusNoContent)
				return nil
			}
			return next(req)
		}
	}
}
