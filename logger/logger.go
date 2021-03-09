package logger

import (
	"log"
	"os"
)

const logFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile

var (
	Error = log.New(os.Stdout, "ERROR ", logFlags)
	Warn  = log.New(os.Stdout, "WARN ", logFlags)
	Info  = log.New(os.Stdout, "INFO ", logFlags)
)
