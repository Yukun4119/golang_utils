package log

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

const (
	InfoLevel  = Green + "[INFO]" + Reset
	DebugLevel = Yellow + "[DEBUG]" + Reset
	ErrorLevel = Red + "[ERROR]" + Reset
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

const (
	Whitespace = " "
	Newline    = "\n"
)
