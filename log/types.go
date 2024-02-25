package log

import (
	"bytes"
	"io"
	"sync"
)

type Level int32
type Processor func(format string, v ...any) (string, []any)

type Logger struct {
	level        Level
	prefix       string
	fileLocation string
	showDetail   bool
	mutex        sync.Mutex
	buf          bytes.Buffer
	w            io.Writer
	processors   []Processor
}
