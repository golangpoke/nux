package nux_test

import (
	"github.com/golangpoke/nux/code"
	"github.com/golangpoke/nux/nlog"
	"github.com/golangpoke/nux/nux"
	"io"
	"os"
	"path"
	"testing"
)

func TestHello(t *testing.T) {
	defer nlog.Recovery()
	n := nux.New()
	n.Use(nux.Recovery(), nux.Logger())
	api := n.Group("/api")
	v1 := api.Group("/v1")
	v1.POST("/test", HandleTest())
	v1.POST("/upload", UploadTest())
	n.Start(":8000")
}

func UploadTest() nux.HandleFunc {
	return func(req *nux.Request) nux.Response {
		nlog.INFOf("upload test")
		req.HeaderSet("Access-Control-Allow-Origin", "*")

		fileHeaders, err := req.ParseMultiFileHeaders(32<<20, "files")
		if err != nil {
			return code.ErrBadRequest.With(err)
		}
		for _, fileHeader := range fileHeaders {
			fileHeaderOpen, err := fileHeader.Open()
			if err != nil {
				nlog.Panic(err)
			}
			fileCreate, err := os.Create(path.Join("test", fileHeader.Filename))
			if err != nil {
				nlog.Panic(err)
			}
			_, err = io.Copy(fileCreate, fileHeaderOpen)
			if err != nil {
				nlog.Panic(err)
			}
			_ = fileHeaderOpen.Close()
			_ = fileCreate.Close()
		}
		return nil
	}
}

func HandleTest() nux.HandleFunc {
	return func(req *nux.Request) nux.Response {
		nlog.INFOf("handle test")
		data := struct {
			Data string `json:"data"`
		}{}
		err := req.Bind(&data)
		if err != nil {
			return code.ErrBadRequest.With(err)
		}
		// if err := os.ErrNotExist; err != nil {
		// 	nlog.Panic(err)
		// }
		return nux.Map{
			"hello": data.Data,
		}
	}
}
