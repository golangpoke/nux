package nux

import (
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	w http.ResponseWriter
	r *http.Request
}

func newRequest(w http.ResponseWriter, r *http.Request) *Request {
	return &Request{w: w, r: r}
}

func (r *Request) sendJson(code int, data any) {
	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(code)
	encoder := json.NewEncoder(r.w)
	if err := encoder.Encode(data); err != nil {
		http.Error(r.w, err.Error(), http.StatusInternalServerError)
	}
}

func (r *Request) Bind(data any) {
	body, err := io.ReadAll(r.r.Body)
	if err != nil {
		http.Error(r.w, err.Error(), http.StatusBadRequest)
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		http.Error(r.w, err.Error(), http.StatusBadRequest)
	}
}

func (r *Request) Method() string {
	return r.r.Method
}

func (r *Request) Url() string {
	return r.r.URL.Path
}

func (r *Request) PathValue(key string) string {
	return r.r.PathValue(key)
}
