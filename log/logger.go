package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"time"
)

func (l *Logger) SetLevel(level Level) {
	atomic.StoreInt32((*int32)(&l.level), int32(level))
}

func (l *Logger) GetLevel() Level {
	return Level(atomic.LoadInt32((*int32)(&l.level)))
}

func (l *Logger) Info(format string, v ...any) {
	if l.level > LevelInfo {
		return
	}
	l.w.Write([]byte(l.assembleMsg(InfoLevel, format, v...)))
}

func (l *Logger) Debug(format string, v ...any) {
	if l.level > LevelDebug {
		return
	}
	l.w.Write([]byte(l.assembleMsg(DebugLevel, format, v...)))
}

func (l *Logger) Error(format string, v ...any) {
	if l.level > LevelError {
		return
	}
	l.w.Write([]byte(l.assembleMsg(ErrorLevel, format, v...)))
}

func (l *Logger) assembleMsg(logLevel string, format string, v ...any) Message {
	msg := logLevel + " "

	if l.showDetail {
		msg += fmt.Sprintf("%s", time.Now()) + " "
	}

	msg += getFileLocation()
	msg += fmt.Sprintf(format, v...) + " "
	msg += "\n"
	return Message(msg)
}

func getFileLocation() string {
	_, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "unknown file"
		line = 0
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line) + " "
}
