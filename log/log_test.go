package log

import "testing"

func TestInfo(t *testing.T) {
	Info("hello info")
}

func TestDebug(t *testing.T) {
	Debug("hello debug")
}

func TestError(t *testing.T) {
	Error("hello error")
}
