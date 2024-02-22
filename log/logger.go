package log

import "sync/atomic"

func (l *Logger) SetLevel(level Level) {
	//l.mutex.Lock()
	//defer l.mutex.Unlock()
	//l.level = level
	atomic.StoreInt32((*int32)(&l.level), int32(level))
}

func (l *Logger) GetLevel() Level {
	return Level(atomic.LoadInt32((*int32)(&l.level)))
}
