package nux

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

type Request struct {
	Writer http.ResponseWriter
	r      *http.Request
}

func newRequest(w http.ResponseWriter, r *http.Request) *Request {
	return &Request{Writer: w, r: r}
}

func (r *Request) sendJson(code int, data any) {
	r.Writer.Header().Set("Content-Type", "application/json")
	r.Writer.WriteHeader(code)
	encoder := json.NewEncoder(r.Writer)
	if err := encoder.Encode(data); err != nil {
		http.Error(r.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (r *Request) Bind(data any) error {
	body, err := io.ReadAll(r.r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Request) ParseMultiFileHeaders(maxMemory int64, key string) ([]*multipart.FileHeader, error) {
	err := r.r.ParseMultipartForm(maxMemory)
	if err != nil {
		return nil, err
	}
	return r.r.MultipartForm.File[key], nil
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
