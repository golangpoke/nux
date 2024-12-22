package nux_test

import (
	"github.com/golangpoke/nux/nlog"
	"github.com/golangpoke/nux/nux"
	"testing"
)

func TestHello(t *testing.T) {
	defer nlog.Recovery()
	n := nux.New()
	n.Use(nux.Recovery(), nux.Logger())
	api := n.Group("/api")
	v1 := api.Group("/v1")
	v1.POST("/test", HandleTest())
	n.Start(":8000")
}

func HandleTest() nux.HandleFunc {
	return func(req *nux.Request) nux.Response {
		nlog.INFOf("handle test")
		data := struct {
			Data string `json:"data"`
		}{}
		req.Bind(&data)
		// if err := os.ErrNotExist; err != nil {
		// 	nlog.Panic(err)
		// }
		return nux.Map{
			"hello": data.Data,
		}
	}
}
