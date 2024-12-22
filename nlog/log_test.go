package nlog_test

import (
	"github.com/golangpoke/nux/nlog"
	"os"
	"testing"
)

func TestPanic(t *testing.T) {
	defer nlog.Recovery()
	err := os.ErrPermission
	if err != nil {
		nlog.ERROf("err:%v", err)
	}
	err = os.ErrNotExist
	if err != nil {
		nlog.Panic(err)
	}
}
