package log

import (
	"testing"
)

func TestInfo(t *testing.T) {
	p := func(format string, v ...any) (string, []any) {
		return "***:%+v", v
	}
	AddProcessor(p)
	d := 11
	Info("hello info :%d", d)
}

func TestDebug(t *testing.T) {
	Debug("hello debug")
}

func TestError(t *testing.T) {
	Error("hello error")
}
