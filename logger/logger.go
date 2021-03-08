package logger

import (
	"log"
	"os"
)

const logFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile

var (
	Warn  = log.New(os.Stdout, "WARN ", logFlags)
	Info  = log.New(os.Stdout, "INFO ", logFlags)
	Error = log.New(os.Stdout, "ERROR ", logFlags)
)
