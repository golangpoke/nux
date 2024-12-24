package nux

import (
	"net/http"
)

type HandleFunc func(req *Request) Response

func (h HandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := newRequest(w, r)
	res := h(req)
	HandleResponseFunc(req, res)
}

type Middleware func(next HandleFunc) HandleFunc

type nux struct {
	mux         *http.ServeMux
	group       string
	middlewares []Middleware
}

func New() *nux {
	return &nux{
		mux:         http.NewServeMux(),
		middlewares: make([]Middleware, 0),
	}
}

func (n *nux) Group(path string) *nux {
	newNux := &nux{
		mux:         n.mux,
		group:       n.group + path,
		middlewares: n.middlewares,
	}
	return newNux
}

func (n *nux) Use(middleware ...Middleware) {
	n.middlewares = append(n.middlewares, middleware...)
}

func (n *nux) GET(router string, handle HandleFunc) {
	n.handleRouter(http.MethodGet, router, handle)
}
func (n *nux) POST(router string, handle HandleFunc) {
	n.handleRouter(http.MethodPost, router, handle)
}
func (n *nux) PUT(router string, handle HandleFunc) {
	n.handleRouter(http.MethodPut, router, handle)
}
func (n *nux) DELETE(router string, handle HandleFunc) {
	n.handleRouter(http.MethodDelete, router, handle)
}

func (n *nux) All(router string, handle HandleFunc) {
	n.handleRouter("", router, handle)
}

func (n *nux) Start(addr string) error {
	return http.ListenAndServe(addr, n.mux)
}
