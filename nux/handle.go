package nux

import (
	"fmt"
	"net/http"
)

var HandleResponseFunc = func(req *Request, res Response) {
	// Request为空时，跳过处理
	if res == nil {
		return
	}
	m := make(Map)
	statusCode := http.StatusOK
	if res.Error() == nil && res.Code() == Success {
		m["data"] = res.Data()
	} else {
		m["error"] = res.Error().Error()
		statusCode = http.StatusInternalServerError
	}
	m["code"] = res.Code()
	m["message"] = res.Message()
	req.sendJson(statusCode, m)
}

func (n *nux) handleRouter(method, router string, handle HandleFunc) {
	router = fmt.Sprintf("%s %s%s", method, n.group, router)
	handle = n.handleMiddlewares(handle)
	n.mux.Handle(router, handle)
}

func (n *nux) handleMiddlewares(handle HandleFunc) HandleFunc {
	// 反向嵌套中间件
	l := len(n.middlewares) - 1
	for i := l; i >= 0; i-- {
		handle = n.middlewares[i](handle)
	}
	return handle
}
