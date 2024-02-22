package log

import (
	"bytes"
	"io"
	"sync"
)

type Level int32

type Logger struct {
	level        Level
	prefix       string
	fileLocation string
	mutex        sync.Mutex
	buf          bytes.Buffer
	w            io.Writer
}
