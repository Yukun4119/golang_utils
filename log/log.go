package log

import (
	"os"
)

var (
	defaultLogger *Logger
)

func init() {
	defaultLogger = NewLogger()
}

/*
NewLogger

By convention, many programs output their log messages to os.Stderr
instead of os.Stdout for a couple of reasons:

1. Separation of Concerns: It allows for the separation of regular
program output from error messages and log information.
This can be very useful when you’re piping or redirecting output
in the command line.
2. Order of Messages: os.Stderr is unbuffered, while os.Stdout is buffered.
This means that if your program crashes or exits unexpectedly, messages
sent to os.Stdout might not get printed if the buffer isn’t flushed before the program exits.
But messages sent to os.Stderr will always get printed immediately.
*/
func NewLogger() *Logger {
	return &Logger{
		level:      LevelInfo,
		w:          os.Stderr,
		showDetail: false,
	}
}

/*
SetLevel

Enum of different levels:

	Debug: 0
	INFO:  1
	ERROR: 2
*/
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

/*
GetLevel
*/
func GetLevel() Level {
	return defaultLogger.GetLevel()
}

/*
Info
*/
func Info(format string, v ...any) {
	defaultLogger.Info(format, v...)
}

/*
Debug
*/
func Debug(format string, v ...any) {
	defaultLogger.Debug(format, v...)
}

/*
Error
*/
func Error(format string, v ...any) {
	defaultLogger.Error(format, v...)
}

func AddProcessor(p Processor) {
	defaultLogger.AddProcessor(p)
}
