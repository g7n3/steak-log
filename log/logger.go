package log

import (
	"io"
	"os"
	"log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*log.Logger
}

var output *lumberjack.Logger
var flags int

func init() {
	output = &lumberjack.Logger{
		Filename:   "log/mee.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     1, // days
	}
	flags = log.Ldate | log.Ltime | log.Lshortfile
	log.SetOutput(io.MultiWriter(output, os.Stdout))
	log.SetFlags(flags)
}

func New(prefix string) *Logger {
	l := log.New(io.MultiWriter(output, os.Stdout), prefix+": ", flags)
	return &Logger{Logger: l}
}

func Println(v ...interface{}) {
	log.Println(v)
}
