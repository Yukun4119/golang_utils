package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func (l *Logger) assembleMsg(logLevel string, format string, v ...any) string {

	// https://golangnote.com/golang/golang-stringsbuilder-vs-bytesbuffer
	// Both `strings.Builder` and `bytes.Buffer` are used for efficient in Golang.
	// Here I choose strings.Builder

	var msg strings.Builder
	msg.WriteString(logLevel)
	msg.WriteString(Whitespace)

	if l.showDetail {
		msg.WriteString(time.Now().String())
		msg.WriteString(Whitespace)
	}

	msg.WriteString(getFileLocation())
	msg.WriteString(l.getMessage(format, v...))
	msg.WriteString(Whitespace)
	msg.WriteString(Newline)

	return msg.String()
}

func getFileLocation() string {
	_, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "unknown file"
		line = -1
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line) + " "
}

func (l *Logger) getMessage(format string, v ...any) string {
	for _, process := range l.processors {
		format, v = process(format, v...)
	}
	return fmt.Sprintf(format, v...)
}
