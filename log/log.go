package log

import "log"

func init() {
	//log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetPrefix("golang_utils_")
}

func Println(v ...any) {
	log.Println(v)
}

func Print(v ...any) {
	log.Print(v)
}

func Printf(format string, v ...any) {
	log.Printf(format, v)
}
