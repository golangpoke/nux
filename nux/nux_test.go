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
	// n.Use(func(next nux.HandleFunc) nux.HandleFunc {
	// 	return func(req *nux.Request) nux.Response {
	// 		// req.Header("Access-Control-Allow-Origin", "*")
	// 		// req.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 		// req.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
	// 		return next(req)
	// 	}
	// })
	n.Use(nux.CORS(), nux.Recovery(), nux.Logger())
	api := n.Group("/api")

	api.POST("/test", HandleTest())
	api.All("/upload/stage", UploadTest())
	n.Start(":8000")
}

func UploadTest() nux.HandleFunc {
	return func(req *nux.Request) nux.Response {
		nlog.INFOf("upload test")
		req.Writer.Header().Set("Access-Control-Allow-Origin", "*")

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
