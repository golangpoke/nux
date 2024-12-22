package nlog

import (
	"fmt"
	"log"
	"runtime"
)

var DEBUG = true

func loadStack() string {
	if !DEBUG {
		return ""
	}
	_, f, l, _ := runtime.Caller(2)
	return fmt.Sprintf(" %s:%d", f, l)
}

func Panic(err error) {
	s := fmt.Sprintf("PANIC %s%s", err, loadStack())
	panic(s)
}

func Recovery() {
	if err := recover(); err != nil {
		log.Printf("%v", err)
	}
}

func INFOf(msg string, args ...any) {
	msg = fmt.Sprintf("INFO %s%s", msg, loadStack())
	log.Printf(msg, args...)
}

func WARNf(msg string, args ...any) {
	msg = fmt.Sprintf("WARN %s%s", msg, loadStack())
	log.Printf(msg, args...)
}

func ERROf(msg string, args ...any) {
	msg = fmt.Sprintf("ERRO %s%s", msg, loadStack())
	log.Printf(msg, args...)
}
