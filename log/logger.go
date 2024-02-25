package log

import (
	"sync/atomic"
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

func (l *Logger) AddProcessor(p Processor) {
	l.processors = append(l.processors, p)
}
